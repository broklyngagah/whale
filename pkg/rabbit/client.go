package rabbit

import (
	"github.com/streadway/amqp"
	"fmt"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"time"
	"carp.cn/whale/pkg/rpc"
	"encoding/json"
	"sync"
)

type RabbitMQClient struct {
	url       string
	conn      *amqp.Connection
	ch        *amqp.Channel
	rpc       *rpc.RpcHelper
	queues    *sync.Map                          // map[*amqp.Queue]struct{}
	done      chan struct{}                      // 关闭所有消费者
	task      chan *Message                      // 任务队列
	errorChan chan *amqp.Error                   // 错误事件
	snapshot  func(client *RabbitMQClient) error // 保存整个初始化过程， 重连时调用
}

func NewRabbitMQClient(host string, port int, user, pass string, rpc *rpc.RpcHelper) *RabbitMQClient {
	return &RabbitMQClient{
		url:    fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port),
		rpc:    rpc,
		queues: &sync.Map{},
		//done:      make(chan struct{}),
		task:      make(chan *Message, 1024),
		//errorChan: make(chan *amqp.Error),
	}
}

func (r *RabbitMQClient) CreateConnection() error {
	var (
		conn *amqp.Connection
		ch   *amqp.Channel
		err  error
		i    = 0
	)
RECONNECTION:
	i++
	conn, err = amqp.Dial(r.url)
	if err != nil {
		zaplogger.Error("[RabbitMQ]: Connection server error.", zap.Error(err))

		if i > 60 * 3 {
			zaplogger.Error("[RabbitMQ]: Reconnection failed", zap.Error(err))
			return err
		}
		time.Sleep(time.Second)
		goto RECONNECTION
	}
	ch, err = conn.Channel()
	if err != nil {
		return err
	}
	r.conn = conn
	r.ch = ch
	// 重连监听
	r.errorChan = r.conn.NotifyClose(make(chan *amqp.Error))

	r.done = make(chan struct{})
	return nil
}

func (r *RabbitMQClient) listener() {
	defer func() {
		if err := recover() ; err != nil {
			zaplogger.Error("[RabbitMQ]: listener panic", zap.Reflect(":", err))
		}
		// 重连
		r.reconnection()
	}()
	zaplogger.Info("[Rabbit]: -----listen is start-----")
	for {
		select {
		case err, ok := <-r.errorChan:

			if !ok {
				zaplogger.Error("[RabbitMQ]: errorChan is closed")
				return
			}
			if err.Code == amqp.ChannelError || err.Code == amqp.ChannelError {
				close(r.errorChan)
				zaplogger.Error("[RabbitMQ]: error", zap.Int(" Code:", err.Code), zap.String(" Msg:", err.Reason))
				return
			}
			zaplogger.Error("[RabbitMQ]: error", zap.Int(" Code:", err.Code), zap.String(" Msg:", err.Reason))
		case msg, ok := <-r.task:
			// send message
			if !ok {
				zaplogger.Error("RabbitMQ]: task queue is closed")
				return
			}
			err := r.ch.Publish(msg.exchange, msg.key, msg.mandatory, msg.immediate, amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg.data,
			})
			if err != nil {
				zaplogger.Error("[Rabbit]: Publish massage error",
					zap.Error(err), zap.String(" exchange:", msg.exchange), zap.String(" msg:", string(msg.data)))
			}
		case <-r.done:
			return
		}

	}
	zaplogger.Info("[Rabbit]: -----listen is stop-----")
}

func (r *RabbitMQClient) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool) error {
	return r.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, nil)
}

func (r *RabbitMQClient) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool) (*amqp.Queue, error) {
	q, err := r.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	return &q, err
}

func (r *RabbitMQClient) QueueBind(q *amqp.Queue, key, exchange string, noWait bool) error {
	err := r.ch.QueueBind(q.Name, key, exchange, false, nil)
	if err != nil {
		return err
	}
	r.queues.Store(q, struct{}{})
	return nil
}

func (r *RabbitMQClient) Publish(exchange, key string, mandatory, immediate bool, msg []byte) {
	r.task <- NewMessage(exchange, key, mandatory, immediate, msg)
}

func (r *RabbitMQClient) Start() {
	go r.listener()
	go r.Run()
	zaplogger.Info("[RabbitMQ]: ...client start ...")
}

func (r *RabbitMQClient) Run() {
	r.queues.Range(func(key, value interface{}) bool {
		go r.consumer(key.(*amqp.Queue))
		return true
	})
}

func (r *RabbitMQClient) consumer(q *amqp.Queue) {
	msgs, err := r.ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		zaplogger.Error("[RabbitMQ]: create consume error", zap.Error(err))
		return
	}

	defer func() {
		if r := recover(); r != nil {
			zaplogger.Error("[RabbitMQ]: process massage panic", zap.Reflect(":", r))
		}
		zaplogger.Info("[RabbitMQ]: consumer is closed.", zap.String(" queue:", q.Name))
	}()

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				zaplogger.Info("[RabbitMQ]: massage channel is closed.", zap.String(" queue:", q.Name))
				return
			}
			zaplogger.Info("[RabbitMQ]: [IN]", zap.String("Message:", string(msg.Body)))

			var req *rpc.Request
			err := json.Unmarshal(msg.Body, &req)
			if err != nil {
				zaplogger.Error("[RabbitMQ]: json unmarshal error:", zap.Error(err), zap.String(" msg:", string(msg.Body)))
				continue
			}
			resp := r.rpc.Handler(req)
			if resp.Error != nil {
				zaplogger.Error("[RabbitMQ]: process massage error", zap.Reflect(":", resp.Error))
			}

		case <-r.done:
			return
		}
	}
}

func (r *RabbitMQClient) SetSnapshot(f func(client *RabbitMQClient) error) {
	r.snapshot = f
}
func (r *RabbitMQClient) GetSnapshot() func(client *RabbitMQClient) error {
	return r.snapshot
}

func (r *RabbitMQClient) reconnection() {
	zaplogger.Info("[RabbitMQ]: -----Begin Reconnection------")
	r.reset()
	i := 0
RETRY:
	i++

	err := r.CreateConnection()
	if err != nil {
		goto RETRY
	}

	err = r.snapshot(r)
	if err != nil {
		r.ch.Close()
		r.conn.Close()
		goto RETRY
	}
	r.Start()
}

func (r *RabbitMQClient) reset() {

	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("[RabbitMQ]: reset panic", zap.Reflect(":", err))
		}
	}()

	err := r.conn.Close()
	if err != nil {
		zaplogger.Error("[RabbitMQ]: close connection error", zap.Error(err))
	}

	err = r.ch.Close()
	if err != nil {
		zaplogger.Error("[RabbitMQ]: close channel error", zap.Error(err))
	}

	close(r.done)


	r.queues.Range(func(key, value interface{}) bool {
		r.queues.Delete(key)
		return true
	})
}

func (r *RabbitMQClient) Destroy() {
	defer func() {
		if err := recover(); err != nil {
			zaplogger.Error("[RabbitMQ]: destroy panic", zap.Reflect(" :", err))
		}
	}()

	// close all consumer
	close(r.errorChan)
	close(r.task)
	r.reset()
	zaplogger.Info("[RabbitMQ]: client is closed")
}

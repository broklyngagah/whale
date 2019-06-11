package chat

import (
	"carp.cn/whale/pkg/rabbit"
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/config"
	"carp.cn/whale/pkg/rpc"
	"fmt"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"encoding/json"
	"github.com/streadway/amqp"
)

const (
	BroadcastExchange = "whale.chat.broadcast"
	Queue             = "broadcast.chat.node-%d"
	Broadcast         = amqp.ExchangeFanout
)

var DefaultRabbitHelper *rabbit.RabbitMQClient

func getQueueName(node int) string {
	return fmt.Sprintf(Queue, node)
}

func startRabbitClient(c *cli.Context) {
	rpc := rpc.NewRpcHelper()
	rpc.RegisterMethod(NewRabbitLive())
	cfg := config.Get().Amqp
	DefaultRabbitHelper = rabbit.NewRabbitMQClient(cfg.Host, cfg.Port, cfg.User, cfg.Pass, rpc)

	err := DefaultRabbitHelper.CreateConnection()
	if err != nil {
		zaplogger.Fatal("rabbit create connection err", zap.Error(err))
		return
	}

	snapshot := func(r *rabbit.RabbitMQClient) error {
		err := r.ExchangeDeclare(BroadcastExchange, Broadcast, true, false, false, false)
		if err != nil {
			zaplogger.Error("[Rabbit]: Create exchange error ", zap.Error(err))
			return err
		}
		q, err := r.QueueDeclare(getQueueName(c.Int("node")), false, true, false, false)
		if err != nil {
			zaplogger.Error("[Rabbit]: Create queue error ", zap.Error(err))
			return err
		}
		err = r.QueueBind(q, "", BroadcastExchange, false)
		if err != nil {
			zaplogger.Error("[Rabbit]: Queue binding error",
				zap.Error(err), zap.String(" queue:", q.Name), zap.String(" exchange:", BroadcastExchange))
			return err
		}
		return nil
	}
	DefaultRabbitHelper.SetSnapshot(snapshot)
	err = DefaultRabbitHelper.GetSnapshot()(DefaultRabbitHelper)
	if err != nil {
		zaplogger.Error("rabbit run snapshot")
	}
	DefaultRabbitHelper.Start()
}

func broadcastMessage(msg []byte) {
	DefaultRabbitHelper.Publish(BroadcastExchange, "", false, false, msg)
}

func BroadcastChat(funcName string, params ... interface{}) {
	req := &rpc.Request{
		Method: funcName,
		Params: params,
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		zaplogger.Error("BroadcastChat json marshal error", zap.Error(err), zap.Reflect(" req:", req))
		return
	}
	broadcastMessage(bytes)
}

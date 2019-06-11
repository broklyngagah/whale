package rabbit

import (
	"testing"
	"fmt"
	"carp.cn/whale/zaplogger"
	"time"
)

func init(){
	zaplogger.SetLogger("../logs", "test/log", "debug", true)
}

func TestNewRabbitMQClient(t *testing.T) {
	cli := NewRabbitMQClient("www.snlan.top", 5672,"snlan", "123456", nil)
	err := cli.CreateConnection()
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	snap := func(r *RabbitMQClient) error {

		fmt.Println("-------init register-------")
		return nil
	}
	cli.SetSnapshot(snap)
	err = cli.GetSnapshot()(cli)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cli.Start()

	go func() {
		time.Sleep(time.Second*5)
		//cli.errorChan <- &amqp.Error{Code:amqp.ChannelError}
		//cli.conn.Close()
	}()
	select {
	}

}

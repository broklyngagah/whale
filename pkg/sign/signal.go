package sign

import (
	"os"
	"os/signal"
	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
)

type Handler func(chan os.Signal)

type Hook struct {
	os.Signal
	Handler
}

func ListenSignal(hooks ...*Hook){

	var sigList []os.Signal

	for _, h := range hooks {
		sigList = append(sigList, h.Signal)
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, sigList...)

	for sig := range sigChan {
		zaplogger.Info("[CMD] Receive a sign.", zap.Reflect("Signal:", sig))
		for _, h := range hooks {
			if sig == h.Signal {
				h.Handler(sigChan)
			}
		}
	}
}
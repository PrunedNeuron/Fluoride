package config

import (
	"os"
	"os/signal"
	"sync"

	"go.uber.org/zap"
)

type stop struct {
	channel chan struct{}
	sync.WaitGroup
}

var (
	// Stop is the global stop instance
	Stop = &stop{
		channel: make(chan struct{}),
	}

	// Handle signals
	signalChannel = make(chan os.Signal, 1)
)

func init() {

	// Stop flag indicates if ctrl+c has been sent
	signal.Notify(signalChannel, os.Interrupt)

	// Handle signals
	go func() {
		for {
			for sig := range signalChannel {
				switch sig {
				case os.Interrupt:
					zap.S().Info("Received interrupt...")
					close(Stop.channel)
					return
				}
			}
		}
	}()
}

// Chan returns a read only channel that is closed when the program exits
func (s *stop) Chan() <-chan struct{} {
	return s.channel
}

// Bool returns t/f if we should stop
func (s *stop) Bool() bool {
	select {
	case <-s.channel:
		return true
	default:
		return false
	}

}

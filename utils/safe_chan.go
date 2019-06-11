package utils

import (
	"sync"
	"time"
)

type SafeChan struct {
	mu     sync.RWMutex
	closed bool
	ch     chan interface{}
}

func NewSafeChan(size int) *SafeChan {
	if size <= 0 {
		size = 1
	}

	sc := &SafeChan{
		ch: make(chan interface{}, size),
	}
	return sc
}

func (s *SafeChan) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.closed {
		s.closed = true
		close(s.ch)
	}
}

func (s *SafeChan) Closed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.closed
}

func (s *SafeChan) Send(v interface{}) (ok bool) {
	s.mu.Lock()
	ok = false
	defer func() {
		if err := recover(); err != nil {
			//just to recover painc on closed channel
		}
		s.mu.Unlock()
	}()
	s.ch <- v
	ok = true
	return
}

func (s *SafeChan) ReadOnly() <-chan interface{} {
	return s.ch
}

func (s *SafeChan) ReadTimeout(d time.Duration) (data interface{}, ok bool) {
	select {
	case data, ok = <-s.ReadOnly():
		return
	case <-time.After(d):
		ok = false
		data = nil
	}
	return
}

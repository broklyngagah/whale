package chat

import (
	"sync"
	"math"
)

//计数器
type Counter struct {
	id uint32
	mu sync.Mutex
}

var DefaultCounter = NewCounter()

func NewCounter() *Counter {
	return &Counter{
		id: uint32(0),
		mu: sync.Mutex{},
	}
}

func (c *Counter) getId() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.id >= math.MaxUint32 {
		c.id = 0
	}
	c.id++
	return c.id
}

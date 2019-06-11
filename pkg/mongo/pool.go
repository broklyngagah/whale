package mongo

import (
	"gopkg.in/mgo.v2"
)

var Factory *Pool

type Pool struct {
	size int
	idle chan *mgo.Session
}

func NewPool(size int) *Pool {
	return &Pool{size: size, idle: make(chan *mgo.Session, size)}
}

func (p *Pool) Acquire() *mgo.Session {
	select {
	case v := <-p.idle:
		return v
	default:
		return _Session.Copy()
	}
}

func (p *Pool) Release(s *mgo.Session) {
	select {
	case p.idle <- s:
	default:
		s.Close()
	}
}

func (p *Pool) Close() {
	close(p.idle)
}

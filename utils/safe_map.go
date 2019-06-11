package utils

import (
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{}
	sm.m = make(map[interface{}]interface{})
	return sm
}

func (m *SafeMap) Get(k interface{}) interface{} {
	m.mu.RLock()
	val, ok := m.m[k]
	m.mu.RUnlock()

	if ok {
		return val
	}
	return nil
}

func (m *SafeMap) Set(k interface{}, v interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if val, ok := m.m[k]; !ok {
		m.m[k] = v
	} else if val != v {
		m.m[k] = v
	} else {
		return false
	}
	return true
}

func (m *SafeMap) Check(k interface{}) bool {
	m.mu.RLock()
	_, ok := m.m[k]
	m.mu.RUnlock()
	return ok
}

func (m *SafeMap) Delete(k interface{}) {
	m.mu.Lock()
	delete(m.m, k)
	m.mu.Unlock()
}

func (m *SafeMap) Items() map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	m.mu.RLock()
	for k, v := range m.m {
		res[k] = v
	}
	m.mu.RUnlock()
	return res
}

func (m *SafeMap) Count() int {
	m.mu.RLock()
	l := len(m.m)
	m.mu.RUnlock()
	return l
}

func (m *SafeMap) Echo(f func(k interface{}, v interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.m {
		f(k, v)
	}
}

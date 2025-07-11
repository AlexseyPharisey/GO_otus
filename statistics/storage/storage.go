package storage

import "sync"

type Report struct {
	mu sync.RWMutex
	js []byte
}

func (r *Report) Set(b []byte) {
	r.mu.Lock()
	r.js = b
	r.mu.Unlock()
}

func (r *Report) Get() []byte {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.js
}

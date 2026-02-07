package inmemory

import (
	"sync"
)

type InMemoryStorage struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *InMemoryStorage {
	return &InMemoryStorage{
		mu:   sync.RWMutex{},
		data: map[string]string{},
	}
}

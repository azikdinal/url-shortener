package inmemory

import (
	"sync"
)

type InmemoryStorage struct {
	data   sync.Map // id -> fullURL
	nextID int64
}

func New() *InmemoryStorage {
	return &InmemoryStorage{
		data:   sync.Map{},
		nextID: 1,
	}
}

package inmemory

import (
	"context"
	"errors"
)

func (s *InmemoryStorage) GetFullURL(
	ctx context.Context,
	id int64,
) (string, error) {

	value, ok := s.data.Load(id)
	if !ok {
		return "", errNotFound
	}

	fullURL, ok := value.(string)
	if !ok {
		return "", errNotFound
	}

	return fullURL, nil
}

var errNotFound = errors.New("fullURL not found in Memory")
var errIncorrectValue = errors.New("incorrect fullURL value found in memory")

package link

import (
	"errors"
	"hash/fnv"
	"strings"
)

// Для хеширования ссылок используется LinkID и seed

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	base     = int64(len(alphabet))
	length   = 10
)

var seedInt = seedToInt64("sdfjsdsni")

func encodeFullURL(id int64) string {
	value := id ^ seedInt

	buf := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		buf[i] = alphabet[value%base]
		value /= base
	}

	return string(buf)
}

func seedToInt64(seed string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(seed))
	return int64(h.Sum64())
}

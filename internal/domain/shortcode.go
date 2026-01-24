package domain

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type ShortCode string

func NewShortCode(row string) (ShortCode, error) {
	rowLen := len(row)
	if rowLen != 10 {
		return "", fmt.Errorf("ShortCode character amount must be equal 10, but got %d!", rowLen)
	}
	for i := 0; i < 10; i++ {
		c := row[i]
		if (c >= '0' && c <= '9') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') ||
			c == '_' {
			continue
		}
		return "", fmt.Errorf("%c symbol cannot be part of a code", c)
	}
	return ShortCode(row), nil
}

func generateShortCode() ShortCode {
	b := make([]byte, 6)
	rand.Read(b)
	// TODO: нужно доделать эту функцию, код не соответствует требованиям
	return ShortCode(base64.RawURLEncoding.EncodeToString(b))
}

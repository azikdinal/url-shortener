package link

import (
	"strings"
	"testing"
)

// Проверка длины кода
func TestGenerateShortCode_Length(t *testing.T) {
	shortCode := GenerateShortCode(555)

	if len(shortCode) != length {
		t.Fatalf("length = %d; want %d", len(shortCode), length)
	}
}

// Проверка допустимых символов
func TestGenerateShortCode_Alphabet(t *testing.T) {
	shortCode := GenerateShortCode(341)

	for i, c := range shortCode {
		if !strings.ContainsRune(alphabet, c) {
			t.Fatalf("invalid char %q at position %d", c, i)
		}
	}
}

// Проверка на детерминированность (один id -- один shortCode)
func TestGenerateShortCode_Determenistic(t *testing.T) {
	id := int64(52342)

	shortCode1 := GenerateShortCode(id)
	shortCode2 := GenerateShortCode(id)

	if shortCode1 != shortCode2 {
		t.Fatalf(`codes are not determenistic. 
			shortCode1: %s vs shortCode2: %s`,
			shortCode1,
			shortCode2,
		)
	}
}

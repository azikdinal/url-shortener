package link

import (
	"fmt"
	"math"
	"testing"
)

// Проверка связки generate-parse
func TestShortCodeRoadTrip(t *testing.T) {
	ids := []int64{
		1,
		42,
		54321,
		math.MaxInt32,
	}

	for _, id := range ids {
		t.Run(fmt.Sprintf("id=%d", id), func(t *testing.T) {
			shortCode := GenerateShortCode(id)
			parsed := ParseShortCode(shortCode)

			if parsed != id {
				t.Fatalf("round-trip failed: id=%d -> %s -> %d",
					id, shortCode, parsed)
			}
		})
	}
}

// Проверка на детерминированность (один shortCode -- один id)
func TestParseShortCode_Determenistic(t *testing.T) {
	shortCode := GenerateShortCode(12312)

	id1 := ParseShortCode(shortCode)
	id2 := ParseShortCode(shortCode)

	if id1 != id2 {
		t.Fatalf(`ParseShortCode() is not deterministic.
			id1: %d vs id2: %d`,
			id1,
			id2,
		)
	}
}

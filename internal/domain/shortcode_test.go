package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewShortCode(t *testing.T) {
	val := "aaadf3234_"
	sc, err := NewShortCode(val)
	require.NoError(t, err)
	require.Equal(t, ShortCode(val), sc)
}

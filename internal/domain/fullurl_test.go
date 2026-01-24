package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFullURL_Invalid(t *testing.T) {
	_, err := NewFullURL("noturl")
	require.Error(t, err)
}

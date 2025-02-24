package memstorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacters_ReturnsCharacterStorage(t *testing.T) {
	s := NewStorage()

	count := s.Characters().Count(context.Background())

	assert.Equal(t, 0, count)
}

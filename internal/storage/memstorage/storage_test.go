package memstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacters_ReturnsCharacterStorage(t *testing.T) {
	s := NewStorage()

	assert.NotNil(t, s.Characters())
}

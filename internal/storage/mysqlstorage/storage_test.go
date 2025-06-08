package mysqlstorage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCharacters_ReturnsCharacterStorage(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()

	s := NewStorage(db)

	assert.NotNil(t, s.Characters())
}

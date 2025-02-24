package server

import (
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestServer_NewServer(t *testing.T) {
	mockedStorage := storage.NewMockStorage(t)

	server := NewServer(mockedStorage)

	assert.NotNil(t, server)
}

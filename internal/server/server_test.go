package server

import (
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestServer_NewServer(t *testing.T) {
	mockedStorage := mocks.NewMockStorage(t)

	server := NewServer(mockedStorage)

	assert.NotNil(t, server)
}

package server

import (
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer_NewServer(t *testing.T) {
	mockedStorage := storage.NewMockStorage(t)
	mockedStorage.EXPECT().ImportData(mock.Anything).Once()

	server := NewServer(mockedStorage)

	assert.NotNil(t, server)
}

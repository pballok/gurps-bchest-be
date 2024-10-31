package storage_mem

import "github.com/pballok/gurps-bchest-be/internal/storage"

func NewStorage() storage.Storage {
	return storage.Storage{
		Characters: newCharacterStorage(),
	}
}

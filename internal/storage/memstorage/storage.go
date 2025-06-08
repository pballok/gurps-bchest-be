package memstorage

import (
	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type memStorage struct {
	characters storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType]
}

func (s *memStorage) Characters() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return s.characters
}

func NewStorage() storage.Storage {
	s := memStorage{
		characters: NewCharacterStorable(),
	}
	return &s
}

package storage_mem

import (
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStore struct {
	characters map[storage.CharacterKeyType]*character.Character
}

func newCharacterStorage() storage.Storable[storage.CharacterKeyType, character.Character] {
	return &characterStore{}
}

func (s *characterStore) Add(chr character.Character) (storage.CharacterKeyType, error) {
	id := storage.CharacterKeyType{
		Name:     chr.Name(),
		Campaign: chr.Campaign(),
	}

	_, exists := s.characters[id]
	if exists {
		return storage.CharacterKeyType{}, fmt.Errorf("conflict error: Character with name %s and campaign %s already exists", chr.Name(), chr.Campaign())
	}

	s.characters[id] = &chr
	return id, nil
}

func (*characterStore) Update(id storage.CharacterKeyType, character character.Character) error {
	return nil
}

func (*characterStore) Delete(id storage.CharacterKeyType) error {
	return nil
}

func (*characterStore) Get(id storage.CharacterKeyType) (character.Character, error) {
	return character.NewCharacter("Test", "NPC", "Campaign", 100), nil
}

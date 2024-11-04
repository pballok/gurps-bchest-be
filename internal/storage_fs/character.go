package storage_fs

import (
	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStore struct {
}

func newCharacterStorage() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return &characterStore{}
}

func (*characterStore) Add(chr character.Character) (storage.CharacterKeyType, error) {
	return storage.CharacterKeyType{}, nil
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

func (*characterStore) List(filters storage.CharacterFilterType) ([]character.Character, error) {
	return nil, nil
}

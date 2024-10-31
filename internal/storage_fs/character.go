package storage_fs

import (
	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStore struct {
}

func newCharacterStorage() storage.Storable[character.Character] {
	return &characterStore{}
}

func (*characterStore) Add(character character.Character) (string, error) {
	return "", nil
}

func (*characterStore) Update(id string, character character.Character) error {
	return nil
}

func (*characterStore) Delete(id string) error {
	return nil
}

func (*characterStore) Get(id string) (character.Character, error) {
	return character.NewCharacter("Test", "NPC", "Campaign", 100), nil
}

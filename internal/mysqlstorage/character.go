package mysqlstorage

import (
	"context"
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStorable struct {
}

func NewCharacterStorable() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return &characterStorable{}
}

func (s *characterStorable) Add(_ context.Context, chr character.Character) (storage.CharacterKeyType, error) {
	return storage.CharacterKeyType{}, fmt.Errorf(`failed to add character with campaign "%s", name "%s"`, chr.Campaign(), chr.Name())
}

func (*characterStorable) Update(_ context.Context, id storage.CharacterKeyType, character character.Character) error {
	return nil
}

func (*characterStorable) Delete(_ context.Context, id storage.CharacterKeyType) error {
	return nil
}

func (s *characterStorable) Count(_ context.Context) int {
	return 0
}

func (s *characterStorable) Get(_ context.Context, id storage.CharacterKeyType) (character.Character, error) {
	return nil, fmt.Errorf(`character with campaign "%s", name "%s" not found`, id.Campaign, id.Name)
}

func (s *characterStorable) List(_ context.Context, filters storage.CharacterFilterType) []character.Character {
	return []character.Character{}
}

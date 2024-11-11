package memstorage

import (
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStorable struct {
	characters map[storage.CharacterKeyType]character.Character
}

func NewCharacterStorable() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return &characterStorable{
		characters: make(map[storage.CharacterKeyType]character.Character),
	}
}

func (s *characterStorable) Add(chr character.Character) (storage.CharacterKeyType, error) {
	id := storage.CharacterKeyType{
		Name:     chr.Name(),
		Campaign: chr.Campaign(),
	}

	_, exists := s.characters[id]
	if exists {
		return storage.CharacterKeyType{}, fmt.Errorf(`conflict error: Character with name "%s" and campaign "%s" already exists`, chr.Name(), chr.Campaign())
	}

	s.characters[id] = chr
	return id, nil
}

func (*characterStorable) Update(id storage.CharacterKeyType, character character.Character) error {
	return nil
}

func (*characterStorable) Delete(id storage.CharacterKeyType) error {
	return nil
}

func (s *characterStorable) Count() int {
	return len(s.characters)
}

func (s *characterStorable) Get(id storage.CharacterKeyType) (character.Character, error) {
	c, exists := s.characters[id]
	if !exists {
		return nil, fmt.Errorf(`character with campaign "%s", name "%s" not found`, id.Campaign, id.Name)
	}

	return c, nil
}

func (s *characterStorable) List(filters storage.CharacterFilterType) []character.Character {
	chars := make([]character.Character, 0)
	for _, c := range s.characters {
		if filters.Campaign != nil && *(filters.Campaign) == c.Campaign() {
			chars = append(chars, c)
		}
	}

	return chars
}

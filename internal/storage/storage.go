package storage

import (
	"github.com/pballok/gurps-bchest-be/internal/character"
)

type Storage interface {
	Characters() Storable[CharacterKeyType, character.Character, CharacterFilterType]
}

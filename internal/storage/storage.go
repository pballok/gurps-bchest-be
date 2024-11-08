package storage

import "github.com/pballok/gurps-bchest-be/internal/character"

type Storable[K comparable, V any, F any] interface {
	Add(item V) (K, error)
	Update(id K, item V) error
	Delete(id K) error
	Count() int
	Get(id K) (V, error)
	List(filters F) []V
}

type CharacterKeyType struct {
	Name     string
	Campaign string
}

type CharacterFilterType struct {
	Campaign *string
}

type Storage struct {
	Characters Storable[CharacterKeyType, character.Character, CharacterFilterType]
}

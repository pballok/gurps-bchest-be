package storage

import "github.com/pballok/gurps-bchest-be/internal/character"

type Storable[K comparable, V any] interface {
	Add(item V) (K, error)
	Update(id K, item V) error
	Delete(id K) error
	Get(id K) (V, error)
}

type CharacterKeyType struct {
	Name     string
	Campaign string
}

type Storage struct {
	Characters Storable[CharacterKeyType, character.Character]
}

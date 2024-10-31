package storage

import "github.com/pballok/gurps-bchest-be/internal/character"

type Storable[S any] interface {
	Add(item S) (string, error)
	Update(id string, item S) error
	Delete(id string) error
	Get(id string) (S, error)
}

type Storage struct {
	Characters Storable[character.Character]
}

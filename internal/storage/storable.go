package storage

import "context"

type Storable[K comparable, V any, F any] interface {
	Add(ctx context.Context, item V) (K, error)
	Update(ctx context.Context, id K, item V) error
	Delete(ctx context.Context, id K) error
	Get(ctx context.Context, id K) (V, error)
	List(ctx context.Context, filters F) ([]V, error)
}

type CharacterKeyType struct {
	Name     string
	Campaign string
}

type CharacterFilterType struct {
	Campaign *string
}

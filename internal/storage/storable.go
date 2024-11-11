package storage

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

package mysqlstorage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type characterStorable struct {
	store *sql.DB
}

func NewCharacterStorable(db *sql.DB) storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return &characterStorable{
		store: db,
	}
}

func (s *characterStorable) Add(ctx context.Context, chr character.Character) (storage.CharacterKeyType, error) {
	_, err := s.store.ExecContext(
		ctx,
		"INSERT INTO `character` (name, campaign, player, points) VALUES (?, ?, ?, ?)",
		chr.Name(),
		chr.Campaign(),
		chr.Player(),
		chr.Points(),
	)
	if err != nil {
		return storage.CharacterKeyType{}, err
	}

	return storage.CharacterKeyType{
		Name:     chr.Name(),
		Campaign: chr.Campaign(),
	}, nil
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

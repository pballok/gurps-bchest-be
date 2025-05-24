package mysqlstorage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
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
	query :=
		"INSERT INTO `character` " +
			"(name, campaign, player, points, st_modif, dx_modif, iq_modif, ht_modif, hp_modif, currhp_modif, will_modif, per_modif, fp_modif, currfp_modif, bs_modif, bm_modif)" +
			"  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := s.store.ExecContext(
		ctx,
		query,
		chr.Name(),
		chr.Campaign(),
		chr.Player(),
		chr.Points(),
		chr.Attribute(model.AttributeTypeSt).Modifier(),
		chr.Attribute(model.AttributeTypeDx).Modifier(),
		chr.Attribute(model.AttributeTypeIq).Modifier(),
		chr.Attribute(model.AttributeTypeHt).Modifier(),
		chr.Attribute(model.AttributeTypeHp).Modifier(),
		chr.Attribute(model.AttributeTypeCurrHp).Modifier(),
		chr.Attribute(model.AttributeTypeWill).Modifier(),
		chr.Attribute(model.AttributeTypePer).Modifier(),
		chr.Attribute(model.AttributeTypeFp).Modifier(),
		chr.Attribute(model.AttributeTypeCurrFp).Modifier(),
		chr.Attribute(model.AttributeTypeBs).Modifier(),
		chr.Attribute(model.AttributeTypeBm).Modifier(),
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

func (s *characterStorable) Get(_ context.Context, id storage.CharacterKeyType) (character.Character, error) {
	return nil, fmt.Errorf(`character with campaign "%s", name "%s" not found`, id.Campaign, id.Name)
}

func (s *characterStorable) List(_ context.Context, filters storage.CharacterFilterType) []character.Character {
	return []character.Character{}
}

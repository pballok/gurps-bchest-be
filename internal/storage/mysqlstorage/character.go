package mysqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (*characterStorable) Update(_ context.Context, _ storage.CharacterKeyType, _ character.Character) error {
	return nil
}

func (*characterStorable) Delete(_ context.Context, _ storage.CharacterKeyType) error {
	return nil
}

func (s *characterStorable) Get(ctx context.Context, id storage.CharacterKeyType) (character.Character, error) {
	query :=
		"SELECT name, campaign, player, points, st_modif, dx_modif, iq_modif, ht_modif, hp_modif, currhp_modif, will_modif, per_modif, fp_modif, currfp_modif, bs_modif, bm_modif" +
			" FROM `character` WHERE name=? AND campaign=?"

	var name, campaign, player string
	var points int
	var stModif, dxModif, iqModif, htModif, hpModif, currHpModif, willModif, perModif, fpModif, currFpModif, bsModif, bmModif float64
	err := s.store.QueryRowContext(
		ctx, query, id.Name, id.Campaign,
	).Scan(&name, &campaign, &player, &points, &stModif, &dxModif, &iqModif, &htModif, &hpModif, &currHpModif, &willModif, &perModif, &fpModif, &currFpModif, &bsModif, &bmModif)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf(`character with capaign "%s", name "%s" not found`, id.Campaign, id.Name)
	}
	if err != nil {
		return nil, err
	}

	c := character.NewCharacter(name, player, campaign, points)
	c.Attribute(model.AttributeTypeSt).SetModifier(stModif)
	c.Attribute(model.AttributeTypeDx).SetModifier(dxModif)
	c.Attribute(model.AttributeTypeIq).SetModifier(iqModif)
	c.Attribute(model.AttributeTypeHt).SetModifier(htModif)
	c.Attribute(model.AttributeTypeHp).SetModifier(hpModif)
	c.Attribute(model.AttributeTypeCurrHp).SetModifier(currHpModif)
	c.Attribute(model.AttributeTypeWill).SetModifier(willModif)
	c.Attribute(model.AttributeTypePer).SetModifier(perModif)
	c.Attribute(model.AttributeTypeFp).SetModifier(fpModif)
	c.Attribute(model.AttributeTypeCurrFp).SetModifier(currFpModif)
	c.Attribute(model.AttributeTypeBs).SetModifier(bsModif)
	c.Attribute(model.AttributeTypeBm).SetModifier(bmModif)

	return c, nil
}

func (s *characterStorable) List(ctx context.Context, filters storage.CharacterFilterType) ([]character.Character, error) {
	query :=
		"SELECT name, campaign, player, points, st_modif, dx_modif, iq_modif, ht_modif, hp_modif, currhp_modif, will_modif, per_modif, fp_modif, currfp_modif, bs_modif, bm_modif" +
			" FROM `character`"
	conditions := make([]string, 0)
	if filters.Campaign != nil {
		conditions = append(conditions, "campaign=\""+*filters.Campaign+"\"")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := s.store.QueryContext(ctx, query)
	if err != nil {
		return []character.Character{}, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var name, campaign, player string
	var points int
	var stModif, dxModif, iqModif, htModif, hpModif, currHpModif, willModif, perModif, fpModif, currFpModif, bsModif, bmModif float64
	characters := make([]character.Character, 0)
	for rows.Next() {
		err = rows.Scan(&name, &campaign, &player, &points, &stModif, &dxModif, &iqModif, &htModif, &hpModif, &currHpModif, &willModif, &perModif, &fpModif, &currFpModif, &bsModif, &bmModif)
		if err != nil {
			return []character.Character{}, err
		}

		c := character.NewCharacter(name, player, campaign, points)
		c.Attribute(model.AttributeTypeSt).SetModifier(stModif)
		c.Attribute(model.AttributeTypeDx).SetModifier(dxModif)
		c.Attribute(model.AttributeTypeIq).SetModifier(iqModif)
		c.Attribute(model.AttributeTypeHt).SetModifier(htModif)
		c.Attribute(model.AttributeTypeHp).SetModifier(hpModif)
		c.Attribute(model.AttributeTypeCurrHp).SetModifier(currHpModif)
		c.Attribute(model.AttributeTypeWill).SetModifier(willModif)
		c.Attribute(model.AttributeTypePer).SetModifier(perModif)
		c.Attribute(model.AttributeTypeFp).SetModifier(fpModif)
		c.Attribute(model.AttributeTypeCurrFp).SetModifier(currFpModif)
		c.Attribute(model.AttributeTypeBs).SetModifier(bsModif)
		c.Attribute(model.AttributeTypeBm).SetModifier(bmModif)
		characters = append(characters, c)
	}

	return characters, nil
}

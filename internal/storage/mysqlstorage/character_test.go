package mysqlstorage

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
	"github.com/stretchr/testify/assert"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

func TestCharacterStorable_NewStore(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	assert.NotNil(t, 0, s)
}

func TestCharacterStorable_Add_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func() { _ = db.Close() }()

	c := character.NewCharacter("Test", "Player", "Campaign", 100)

	mock.ExpectExec("INSERT INTO `character`").WithArgs(
		c.Name(), c.Campaign(), c.Player(), c.Points(), 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	s := NewCharacterStorable(db)

	key, err := s.Add(context.Background(), c)

	assert.NoError(t, err)
	assert.Equal(t, "Test", key.Name)
	assert.Equal(t, "Campaign", key.Campaign)
}

func TestCharacterStorable_Add_DBError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func() { _ = db.Close() }()

	c := character.NewCharacter("Test", "Player", "Campaign", 100)

	mock.ExpectExec("INSERT INTO `character`").WillReturnError(sql.ErrConnDone)

	s := NewCharacterStorable(db)

	_, err := s.Add(context.Background(), c)

	assert.Error(t, err)
}

func TestCharacterStorable_Update_Success(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	c := character.NewCharacter("Test", "Player", "Campaign", 10)

	err := s.Update(context.Background(), storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"}, c)

	assert.Nil(t, err)
}

func TestCharacterStorable_Delete_Success(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	err := s.Delete(context.Background(), storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"})

	assert.Nil(t, err)
}

func TestCharacterStorable_Get_Success(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	id := storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"}
	_, err := s.Get(context.Background(), id)

	assert.Error(t, err)
}

func TestCharacterStorable_List_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func() { _ = db.Close() }()

	campaign := "Campaign1"
	columns := []string{"name", "campaign", "player", "points", "st_modif", "dx_modif", "iq_modif", "ht_modif",
		"hp_modif", "currhp_modif", "will_modif", "per_modif", "fp_modif", "currfp_modif", "bs_modif", "bm_modif"}

	mock.ExpectQuery(
		"SELECT " + strings.Join(columns, ", ") + " FROM `character`" +
			" WHERE campaign=\"" + campaign + "\"",
	).WillReturnRows(sqlmock.NewRows(columns).AddRow(
		"Test1", campaign, "Player1", 120, 1, 1, 1, 1, 2, -1, 2, 2, 2, -1, 0, 0,
	).AddRow(
		"Test2", campaign, "Player2", 120, 2, 2, 2, 2, 3, 0, 3, 3, 3, 0, 0, 0,
	))

	s := NewCharacterStorable(db)

	filtered, err := s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})

	assert.NoError(t, err)
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, "Test1", filtered[0].Name())
	assert.Equal(t, "Player1", filtered[0].Player())
	assert.Equal(t, campaign, filtered[0].Campaign())
	assert.Equal(t, 11.0, filtered[0].Attribute(model.AttributeTypeSt).Value())
	assert.Equal(t, "Test2", filtered[1].Name())
	assert.Equal(t, "Player2", filtered[1].Player())
	assert.Equal(t, campaign, filtered[1].Campaign())
	assert.Equal(t, 12.0, filtered[1].Attribute(model.AttributeTypeSt).Value())
}

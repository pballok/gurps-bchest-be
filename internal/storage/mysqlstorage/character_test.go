package mysqlstorage

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

func TestCharacterStorable_NewStore(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	assert.NotNil(t, 0, s)
	assert.Equal(t, 0, s.Count(context.Background()))
}

func TestCharacterStorable_Add_Success(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	c := character.NewCharacter("Test", "Player", "Campaign", 10)
	_, err := s.Add(context.Background(), c)

	assert.NoError(t, err)
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

func TestCharacterStorable_Count_EmptyStorage(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	assert.Equal(t, 0, s.Count(context.Background()))
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
	db, _, _ := sqlmock.New()
	defer func() { _ = db.Close() }()
	s := NewCharacterStorable(db)

	campaign := "Campaign1"
	filtered := s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})
	assert.Equal(t, 0, len(filtered))
}

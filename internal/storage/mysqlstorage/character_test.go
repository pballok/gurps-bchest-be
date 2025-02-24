package mysqlstorage

import (
	"context"
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestCharacterStorable_NewStore(t *testing.T) {
	s := NewCharacterStorable()

	assert.NotNil(t, 0, s)
	assert.Equal(t, 0, s.Count(context.Background()))
}

func TestCharacterStorable_Add_Success(t *testing.T) {
	s := NewCharacterStorable()
	c := character.NewCharacter("Test", "Player", "Campaign", 10)
	_, err := s.Add(context.Background(), c)

	assert.Error(t, err)
}

func TestCharacterStorable_Update_Success(t *testing.T) {
	s := NewCharacterStorable()
	c := character.NewCharacter("Test", "Player", "Campaign", 10)

	err := s.Update(context.Background(), storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"}, c)

	assert.Nil(t, err)
}

func TestCharacterStorable_Delete_Success(t *testing.T) {
	s := NewCharacterStorable()

	err := s.Delete(context.Background(), storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"})

	assert.Nil(t, err)
}

func TestCharacterStorable_Count_EmptyStorage(t *testing.T) {
	s := NewCharacterStorable()
	assert.Equal(t, 0, s.Count(context.Background()))
}

func TestCharacterStorable_Get_Success(t *testing.T) {
	s := NewCharacterStorable()
	id := storage.CharacterKeyType{Name: "Test", Campaign: "Campaign"}
	_, err := s.Get(context.Background(), id)

	assert.Error(t, err)
}

func TestCharacterStorable_List_Success(t *testing.T) {
	s := NewCharacterStorable()

	campaign := "Campaign1"
	filtered := s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})
	assert.Equal(t, 0, len(filtered))
}

package memstorage

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
	id, err := s.Add(context.Background(), c)

	assert.Nil(t, err)
	assert.Equal(t, "Test", id.Name)
	assert.Equal(t, "Campaign", id.Campaign)
	assert.Equal(t, 1, s.Count(context.Background()))
}

func TestCharacterStorable_Add_Fail(t *testing.T) {
	s := NewCharacterStorable()
	c := character.NewCharacter("Test", "Player", "Campaign", 10)
	_, _ = s.Add(context.Background(), c)
	_, err := s.Add(context.Background(), c) // Add same character again
	assert.NotNil(t, err)
	assert.Equal(t, 1, s.Count(context.Background()))
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

func TestCharacterStorable_Count_StorageWithItems(t *testing.T) {
	s := NewCharacterStorable()
	_, _ = s.Add(context.Background(), character.NewCharacter("Test1", "Player1", "Campaign", 10))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test2", "Player2", "Campaign", 10))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test3", "Player3", "Campaign", 10))

	assert.Equal(t, 3, s.Count(context.Background()))
}

func TestCharacterStorable_Get_Success(t *testing.T) {
	s := NewCharacterStorable()
	c := character.NewCharacter("Test", "Player", "Campaign", 10)
	id, _ := s.Add(context.Background(), c)
	addedChar, err := s.Get(context.Background(), id)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Test", addedChar.Name())
}

func TestCharacterStorable_Get_ReturnsWithExpectedItem(t *testing.T) {
	s := NewCharacterStorable()
	_, _ = s.Add(context.Background(), character.NewCharacter("Test1", "Player1", "Campaign", 10))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test2", "Player2", "Campaign", 20))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test3", "Player3", "Campaign", 30))
	addedChar, err := s.Get(context.Background(), storage.CharacterKeyType{
		Name:     "Test2",
		Campaign: "Campaign",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "Test2", addedChar.Name())
	assert.Equal(t, "Player2", addedChar.Player())
	assert.Equal(t, "Campaign", addedChar.Campaign())
	assert.Equal(t, 20, addedChar.Points())
}

func TestCharacterStorable_Get_Fail(t *testing.T) {
	s := NewCharacterStorable()
	_, _ = s.Add(context.Background(), character.NewCharacter("Test", "Player", "Campaign", 10))
	_, err := s.Get(context.Background(), storage.CharacterKeyType{Name: "Test1", Campaign: "Campaign"})

	assert.NotNil(t, err)
}

func TestCharacterStorable_List_Success(t *testing.T) {
	s := NewCharacterStorable()
	_, _ = s.Add(context.Background(), character.NewCharacter("Test1", "Player1", "Campaign1", 10))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test1", "Player1", "Campaign2", 20))
	_, _ = s.Add(context.Background(), character.NewCharacter("Test3", "Player3", "Campaign1", 30))

	campaign := "Campaign1"
	filtered := s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})
	assert.Equal(t, 2, len(filtered))

	campaign = "Campaign2"
	filtered = s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})
	assert.Equal(t, 1, len(filtered))

	campaign = "None"
	filtered = s.List(context.Background(), storage.CharacterFilterType{Campaign: &campaign})
	assert.Equal(t, 0, len(filtered))
}

package storage

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fakeFileInfo struct {
	name string
}

func (f fakeFileInfo) Name() string       { return f.name }
func (f fakeFileInfo) Size() int64        { return 0 }
func (f fakeFileInfo) Mode() os.FileMode  { return os.ModeAppend }
func (f fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (f fakeFileInfo) IsDir() bool        { return false }
func (f fakeFileInfo) Sys() any           { return nil }

type fakeDirEntry struct {
	name string
}

func (f fakeDirEntry) Name() string               { return f.name }
func (f fakeDirEntry) IsDir() bool                { return false }
func (f fakeDirEntry) Type() os.FileMode          { return 0 }
func (f fakeDirEntry) Info() (os.FileInfo, error) { return fakeFileInfo{name: f.name}, nil }

func TestStorage_ImportData_ImportAllCharacters(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "data_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("{\"CharacterName\": \"Character 1\"}"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.EXPECT().Add(mock.Anything, mock.Anything).Times(2).Return(CharacterKeyType{}, nil)

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_ImportData_ImportDirEmpty(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{}, nil)
	mockedFS.AssertNotCalled(t, "ReadFile", mock.Anything)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.AssertNotCalled(t, "Add", mock.Anything)

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_ImportData_ImportDirDoesntExist(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return(nil, os.ErrNotExist)
	mockedFS.AssertNotCalled(t, "ReadFile", mock.Anything)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.AssertNotCalled(t, "Add", mock.Anything)

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_ImportData_ImportFileOpenError(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return(nil, os.ErrPermission)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.EXPECT().Add(mock.Anything, mock.Anything).Once().Return(CharacterKeyType{}, nil)

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_ImportData_InvalidData(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("invalid json"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.EXPECT().Add(mock.Anything, mock.Anything).Once().Return(CharacterKeyType{}, nil)

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_ImportData_AddCharacterError(t *testing.T) {
	mockedFS := NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("{\"CharacterName\": \"Character 1\"}"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.EXPECT().Add(mock.Anything, mock.Anything).Twice().Return(CharacterKeyType{}, fmt.Errorf("error"))

	storageFS = mockedFS
	s := NewStorage(mockedCharacterStorage)

	s.ImportData("./import")
}

func TestStorage_NewStorage_Success(t *testing.T) {
	mockedCharacterStorage := NewMockStorable[CharacterKeyType, character.Character, CharacterFilterType](t)
	mockedCharacterStorage.EXPECT().Count(mock.Anything).Once().Return(42)
	s := NewStorage(mockedCharacterStorage)

	assert.Equal(t, 42, s.Characters().Count(context.Background()))
}

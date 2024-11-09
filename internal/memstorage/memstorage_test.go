package memstorage

import (
	"os"
	"testing"
	"time"

	"github.com/pballok/gurps-bchest-be/internal/mocks/memstorage"
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

func Test_MemStorage_ImportData_ImportDirEmpty(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{}, nil)
	mockedFS.AssertNotCalled(t, "ReadFile", mock.Anything)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 0, s.Characters.Count())
}

func Test_MemStorage_ImportData_ImportAllCharacters(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "data_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("{\"CharacterName\": \"Character 1\"}"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 2, s.Characters.Count())
}

func Test_MemStorage_ImportData_ImportDirDoesntExist(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return(nil, os.ErrNotExist)
	mockedFS.AssertNotCalled(t, "ReadFile", mock.Anything)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 0, s.Characters.Count())
}

func Test_MemStorage_ImportData_ImportFileOpenError(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return(nil, os.ErrPermission)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 1, s.Characters.Count())
}

func Test_MemStorage_ImportData_InvalidData(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("invalid json"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 2\"}"), nil)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 1, s.Characters.Count())
}

func Test_MemStorage_ImportData_ImportSameCharacterTwice(t *testing.T) {
	mockedFS := mocks.NewMockfileSystem(t)
	mockedFS.EXPECT().ReadDir(mock.Anything).Once().Return([]os.DirEntry{
		fakeDirEntry{name: "character_1.json"},
		fakeDirEntry{name: "character_2.json"},
		fakeDirEntry{name: "character_3.json"}}, nil)
	mockedFS.EXPECT().ReadFile("./import/character_1.json").Once().Return([]byte("{\"CharacterName\": \"Character 1\"}"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_2.json").Once().Return([]byte("{\"CharacterName\": \"Character 1\"}"), nil)
	mockedFS.EXPECT().ReadFile("./import/character_3.json").Once().Return([]byte("{\"CharacterName\": \"Character 3\"}"), nil)

	storageFS = mockedFS

	s := New()

	assert.Equal(t, 2, s.Characters.Count())
}

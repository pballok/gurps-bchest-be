package storage

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/pballok/gurps-bchest-be/internal/character"
)

type Storage interface {
	ImportData(string)
	Characters() Storable[CharacterKeyType, character.Character, CharacterFilterType]
}

type storage struct {
	characters Storable[CharacterKeyType, character.Character, CharacterFilterType]
}

func (s *storage) ImportData(importPath string) {
	dataFiles, err := storageFS.ReadDir(importPath)
	if err != nil {
		slog.Any("error", fmt.Errorf("error while checking for import data files: %w", err))
	}
	for _, dataFile := range dataFiles {
		if !dataFile.IsDir() {
			fileName := importPath + "/" + dataFile.Name()
			if strings.HasPrefix(dataFile.Name(), "character_") {
				data, err := storageFS.ReadFile(fileName)
				if err != nil {
					slog.Any("error", fmt.Errorf(`error reading data file "%s": %w`, fileName, err))
					continue
				}
				id, err := s.importCharacter(data)
				if err != nil {
					slog.Any("error", fmt.Errorf(`error importing data file "%s": %w`, fileName, err))
					continue
				}
				slog.Info(fmt.Sprintf(`Imported character "%s" from file "%s"`, id.Name, dataFile.Name()))
			}
		}
	}
}

func (s *storage) importCharacter(data []byte) (CharacterKeyType, error) {
	char, err := character.FromGCA5Import("Imported Campaign", data)
	if err != nil {
		return CharacterKeyType{}, fmt.Errorf(`error importing character: %w`, err)
	}
	id, err := s.characters.Add(char)
	if err != nil {
		return CharacterKeyType{}, fmt.Errorf(`error importing character "%s" to storage: %w`, char.Name(), err)
	}

	return id, nil
}

func (s *storage) Characters() Storable[CharacterKeyType, character.Character, CharacterFilterType] {
	return s.characters
}

func NewStorage(characters Storable[CharacterKeyType, character.Character, CharacterFilterType]) Storage {
	s := storage{
		characters: characters,
	}

	return &s
}

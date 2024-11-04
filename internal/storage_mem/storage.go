package storage_mem

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

func NewStorage() storage.Storage {
	s := storage.Storage{
		Characters: newCharacterStorage(),
	}

	err := importData(&s)
	if err != nil {
		slog.Any("error", err)
	}

	return s
}

func importData(s *storage.Storage) error {
	const importPath = "./import"
	dataFiles, err := os.ReadDir(importPath)
	if err != nil {
		return fmt.Errorf("error while checking for import data files: %w", err)
	}
	for _, dataFile := range dataFiles {
		if !dataFile.IsDir() {
			if strings.HasPrefix(dataFile.Name(), "character_") {
				data, err := os.ReadFile(importPath + "/" + dataFile.Name())
				if err != nil {
					return fmt.Errorf(`error reading data file "%s": %w`, dataFile.Name(), err)
				}
				char, err := character.ImportGCA5Character("Imported Campaign", data)
				if err != nil {
					return fmt.Errorf(`error importing character from file "%s": %w`, dataFile.Name(), err)
				}
				_, err = s.Characters.Add(char)
				if err != nil {
					return fmt.Errorf(`error adding character from file "%s" to storage: %w`, dataFile.Name(), err)
				}

				slog.Info(fmt.Sprintf(`Imported character "%s" from file "%s"`, char.Name(), dataFile.Name()))
			}
		}
	}

	return nil
}

package storage_mem

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type fileSystem interface {
	ReadDir(string) ([]os.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

type osFS struct{}

func (*osFS) ReadDir(path string) ([]os.DirEntry, error) { return os.ReadDir(path) }
func (*osFS) ReadFile(filename string) ([]byte, error)   { return os.ReadFile(filename) }

var storageFS fileSystem = &osFS{}

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
	dataFiles, err := storageFS.ReadDir(importPath)
	if err != nil {
		return fmt.Errorf("error while checking for import data files: %w", err)
	}
	for _, dataFile := range dataFiles {
		if !dataFile.IsDir() {
			if strings.HasPrefix(dataFile.Name(), "character_") {
				data, err := storageFS.ReadFile(importPath + "/" + dataFile.Name())
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

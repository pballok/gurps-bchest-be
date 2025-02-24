package mysqlstorage

import (
	"database/sql"
	"fmt"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/pballok/gurps-bchest-be/internal/utils"
)

type mySQLStorage struct {
	db         *sql.DB
	characters storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType]
}

func (s *mySQLStorage) Characters() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return s.characters
}

func NewDBConnection() (*sql.DB, error) { // coverage-ignore
	dbName := utils.GetEnvOrFail("MYSQL_DATABASE")
	dbUser := utils.GetEnvOrFail("MYSQL_USER")
	dbPwd := utils.GetEnvOrFail("MYSQL_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost)/%s", dbUser, dbPwd, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping mysql: %v", err)
	}

	return db, nil
}

func NewStorage(db *sql.DB) storage.Storage {
	s := mySQLStorage{
		db:         db,
		characters: NewCharacterStorable(),
	}
	return &s
}

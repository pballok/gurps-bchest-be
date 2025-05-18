package mysqlstorage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/pballok/gurps-bchest-be/internal/utils"
)

type mySQLStorage struct {
	characters storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType]
}

func (s *mySQLStorage) Characters() storage.Storable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType] {
	return s.characters
}

func NewDBConnection() (*sql.DB, error) { // coverage-ignore
	dbName := utils.GetEnvOrFail("MYSQL_DATABASE")
	dbUser := utils.GetEnvOrFail("MYSQL_USER")
	dbPwd := utils.GetEnvOrFail("MYSQL_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@tcp(db)/%s", dbUser, dbPwd, dbName)

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

func Migrate(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create mysql driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

func NewStorage(db *sql.DB) storage.Storage {
	s := mySQLStorage{
		characters: NewCharacterStorable(db),
	}
	return &s
}

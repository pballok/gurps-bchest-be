package main

import (
	"log/slog"
	"os"

	"github.com/pballok/gurps-bchest-be/internal/server"
	"github.com/pballok/gurps-bchest-be/internal/storage/mysqlstorage"
)

func configureLogger() {
	logLevel := slog.LevelInfo
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(jsonHandler))
}

func main() {
	configureLogger()

	gurpsDB, err := mysqlstorage.NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer func() { _ = gurpsDB.Close() }()

	err = mysqlstorage.Migrate(gurpsDB)
	if err != nil {
		panic(err)
	}

	gurpsStorage := mysqlstorage.NewStorage(gurpsDB)
	gurpsServer := server.NewServer(gurpsStorage)

	gurpsServer.Run()
	select {}
}

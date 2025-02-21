package main

import (
	"log/slog"
	"os"

	"github.com/pballok/gurps-bchest-be/internal/mysqlstorage"
	"github.com/pballok/gurps-bchest-be/internal/server"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

func configureLogger() {
	logLevel := slog.LevelInfo
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(jsonHandler))
}

func main() {
	configureLogger()

	gurpsStorage := storage.NewStorage(mysqlstorage.NewCharacterStorable())
	gurpsServer := server.NewServer(gurpsStorage)

	gurpsServer.Run()
	select {}
}

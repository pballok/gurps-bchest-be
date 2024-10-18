package main

import (
	"log/slog"
	"os"

	"github.com/pballok/gurps-bchest-be/internal/server"
)

func configureLogger() {
	logLevel := slog.LevelInfo
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(jsonHandler))
}

func main() {
	configureLogger()

	gurpsServer := server.NewServer()
	gurpsServer.Run()
	select {}
}

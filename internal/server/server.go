package server

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pballok/gurps-bchest-be/internal/graph"
	"github.com/pballok/gurps-bchest-be/internal/memstorage"
)

type Server struct {
	server *handler.Server
}

func NewServer() *Server {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Storage: memstorage.New(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/playground"))
	http.Handle("/query", srv)

	return &Server{
		server: srv,
	}
}

func (s *Server) Run() {
	go func() {
		slog.Info("starting server...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			slog.Error("server error: ", slog.Any("error", err))
			os.Exit(1)
		}
	}()
}

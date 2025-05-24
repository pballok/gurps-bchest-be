package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/pballok/gurps-bchest-be/internal/graph"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

type Server struct {
	server *handler.Server
}

func NewServer(storage storage.Storage) *Server {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Storage: storage,
	}}))
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetRecoverFunc(func(ctx context.Context, err any) (userMessage error) {
		slog.Error("graphql server panic: ", slog.Any("error", err))
		return errors.New("graphql server panic")
	})

	http.Handle("/", playground.Handler("GURPS playground", "/playground"))
	http.Handle("/query", srv)

	return &Server{
		server: srv,
	}
}

func (s *Server) Run() { // coverage-ignore
	go func() {
		slog.Info("starting server...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			slog.Error("server error: ", slog.Any("error", err))
			os.Exit(1)
		}
	}()
}

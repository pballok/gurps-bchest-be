package server

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph2 "github.com/pballok/gurps-bchest-be/internal/graph"
)

type Server struct {
	server *handler.Server
}

func NewServer() *Server {
	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: &graph2.Resolver{}}))

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
			slog.Error("server error: ", err)
			os.Exit(1)
		}
	}()
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marceloaguero/go-graphql-products/providers/graph"
	"github.com/marceloaguero/go-graphql-products/providers/model/provider"
	repo "github.com/marceloaguero/go-graphql-products/providers/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbDsn := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	serviceName := os.Getenv("SERVICE_NAME")
	serviceVersion := os.Getenv("SERVICE_VERSION")

	repository, err := repo.NewRepo(dbDsn, dbName)
	if err != nil {
		log.Panic(err)
	}

	usecase := provider.NewUsecase(repository)
	resolver := graph.NewResolver(usecase, serviceName, serviceVersion)

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

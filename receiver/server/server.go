package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sinyavcev/microservices/receiver/graphql/graph"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func Run(db mongo.Database) {
	port := viper.GetString("graphql.port")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Database: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nhe23/aq-api/graph"
	"github.com/nhe23/aq-api/graph/generated"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nhe23/aq-api/pkg/services"
)

const defaultPort = "8080"
const defaultDb = "mongodb://localhost:27018"

// func initDb(dbURI string, dbName string) (db.Collections, error) {
// 	var cols db.Collections
// }

func main() {
	port := os.Getenv("PORT")
	dbURI := os.Getenv("mongodb")
	if port == "" {
		port = defaultPort
	}
	if dbURI == "" {
		dbURI = defaultDb
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal("Could not connect db")
	}
	db := client.Database("AQ_DB")

	locResService := services.NewLocResService(db.Collection("measurements"))
	citiesService := services.NewCitiesService(db.Collection("cities"))
	countriesService := services.NewCountriesService(db.Collection("countries"))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			LocResultsService: locResService,
			CitiesService:     citiesService,
			CountriesSerivce:  countriesService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

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

	"github.com/nhe23/aq-api/services"

	aqdb "github.com/nhe23/aq-api/db"
)

const defaultPort = "8080"

func initDb(dbURI string, dbName string) (aqdb.Collections, error) {
	var cols aqdb.Collections
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
	if err != nil {
		return cols, err
	}
	db := client.Database(dbName)
	cols.CountriesCol = db.Collection("countries")
	cols.CitiesCol = db.Collection("cities")
	cols.MeasurementsCol = db.Collection("measurements")
	return cols, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cols, err := initDb("mongodb://localhost:27018", "AQ_DB")
	if err != nil {
		log.Fatalf("Error initializing collections: %w", err)
	}
	aqdb.SetCollections(&cols)

	locResService := services.NewLocResService()
	citiesService := services.NewCitiesService()
	countriesService := services.NewCountriesService()
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

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nhe23/aq-api/graph"
	"github.com/nhe23/aq-api/graph/generated"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nhe23/aq-api/pkg/services/cities"
	"github.com/nhe23/aq-api/pkg/services/countries"
	"github.com/nhe23/aq-api/pkg/services/measurements"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const defaultPort = "8080"
const defaultDb = "mongodb://localhost:27018"

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "TS:", log.DefaultTimestamp, "caller", log.DefaultCaller)
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
		logger.Log("err", "Error initializing mongo collections")
		os.Exit(1)
	}
	db := client.Database("AQ_DB")

	locResService := measurements.NewService(db.Collection("measurements"))
	locResService = measurements.NewLoggingService(logger, locResService)
	citiesService := cities.NewService(db.Collection("cities"))
	citiesService = cities.NewLoggingService(logger, citiesService)
	countriesService := countries.NewService(db.Collection("countries"))
	countriesService = countries.NewLoggingService(logger, countriesService)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			LocResultsService: locResService,
			CitiesService:     citiesService,
			CountriesSerivce:  countriesService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Printf("connect to http://localhost:%s/ for GraphQL playground\n", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}

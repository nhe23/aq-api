package main

import (
	"context"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/nhe23/aq-api/dataloader"
	"github.com/nhe23/aq-api/graph"
	"github.com/nhe23/aq-api/graph/generated"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nhe23/aq-api/pkg/services/cities"
	"github.com/nhe23/aq-api/pkg/services/countries"
	"github.com/nhe23/aq-api/pkg/services/measurements"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"
const defaultDb = "mongodb://localhost:27018"

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "TS:", log.DefaultTimestamp)
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

	dl := dataloader.NewLoader()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			LocResultsService: locResService,
			CitiesService:     citiesService,
			CountriesSerivce:  countriesService,
			DataLoader:        dl,
		}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost:5000" || r.Host == "localhost:8080"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataloader.Middleware(countriesService, srv))

	logger.Log("connect to http://localhost:%s/ for GraphQL playground\n", port)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

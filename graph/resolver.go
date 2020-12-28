package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/nhe23/aq-api/dataloader"
	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/services/cities"
	"github.com/nhe23/aq-api/pkg/services/countries"
	"github.com/nhe23/aq-api/pkg/services/measurements"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver contains services that are injected into the schema reolver
type Resolver struct {
	LocResultsService measurements.Service
	CitiesService     cities.Service
	CountriesSerivce  countries.Service
	locationResults   []*model.LocationResult
	DataLoader        dataloader.Loader
}

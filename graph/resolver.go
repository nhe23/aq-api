package graph

import "github.com/nhe23/aq-api/pkg/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver contains services that are injected into the schema reolver
type Resolver struct {
	LocResultsService services.LocResService
	CitiesService     services.CitiesSerivce
	CountriesSerivce  services.CountriesSerivce
}

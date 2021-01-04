package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nhe23/aq-api/graph/generated"
	"github.com/nhe23/aq-api/graph/model"
)

func (r *locationResultResolver) Country(ctx context.Context, obj *model.LocationResult) (*model.Country, error) {
	return r.DataLoader.For(ctx).CountryByCode.Load(obj.Country)
}

func (r *queryResolver) Measurements(ctx context.Context, take *int, after *string) ([]*model.LocationResult, error) {
	return r.LocResultsService.GetResults(take, after)
}

func (r *queryResolver) MeasurementsByCountry(ctx context.Context, country string, take *int, after *string) ([]*model.LocationResult, error) {
	return r.LocResultsService.GetResultsByCountry(country, take, after)
}

func (r *queryResolver) MeasurementsByCity(ctx context.Context, city string, take *int, after *string) ([]*model.LocationResult, error) {
	return r.LocResultsService.GetResultsByCity(city, take, after)
}

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	return r.CountriesSerivce.GetCountries()
}

func (r *queryResolver) Cities(ctx context.Context, take *int, after *string) ([]*model.City, error) {
	return r.CitiesService.GetCities(take, after)
}

func (r *queryResolver) CitiesStartsWith(ctx context.Context, searchString string) ([]*model.City, error) {
	return r.CitiesService.CitiesStartsWith(searchString)
}

// LocationResult returns generated.LocationResultResolver implementation.
func (r *Resolver) LocationResult() generated.LocationResultResolver {
	return &locationResultResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type locationResultResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }

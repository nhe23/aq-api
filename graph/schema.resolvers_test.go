package graph

import (
	"context"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/nhe23/aq-api/graph/generated"
	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/mocks"
)

var locResService *mocks.MeasurementsService = new(mocks.MeasurementsService)
var countriesService *mocks.CountriesService = new(mocks.CountriesService)
var citiesService *mocks.CitiesSerivce = new(mocks.CitiesSerivce)
var take int = 1
var after string = "mock"
var c *client.Client = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
	Resolvers: &Resolver{
		LocResultsService: locResService,
		CitiesService:     citiesService,
		CountriesSerivce:  countriesService}})))

func Test_queryResolver_Measurements(t *testing.T) {
	take := 1
	after := "mock"
	location := model.LocationResult{ID: "test"}
	locations := []*model.LocationResult{&location}
	locResService.On("GetResults", &take, &after).Return(locations, nil)
	resolver := &Resolver{LocResultsService: locResService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		take  *int
		after *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LocationResult
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO(), &take, &after}, locations, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Measurements(tt.args.ctx, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.LocationResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.LocationResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_Countries(t *testing.T) {
	country := model.Country{ID: "test", Code: "mock", Name: "testcountry", Count: 5, Cities: 5, Locations: 5}
	countries := []*model.Country{&country}
	countriesService.On("GetCountries").Return(countries, nil)
	resolver := &Resolver{CountriesSerivce: countriesService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Country
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO()}, countries, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Countries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.Countries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.Countries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_Cities(t *testing.T) {
	city := model.City{ID: "test", Name: "MockCity"}
	cities := []*model.City{&city}
	citiesService.On("GetCities", &take, &after).Return(cities, nil)
	resolver := &Resolver{CitiesService: citiesService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		take  *int
		after *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.City
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO(), &take, &after}, cities, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Cities(tt.args.ctx, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.Cities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.Cities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_MeasurementsByCountry(t *testing.T) {
	take := 1
	after := "mock"
	location := model.LocationResult{ID: "test"}
	locations := []*model.LocationResult{&location}
	country := "DE"
	locResService.On("GetResultsByCountry", country, &take, &after).Return(locations, nil)
	resolver := &Resolver{LocResultsService: locResService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx     context.Context
		country string
		take    *int
		after   *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LocationResult
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO(), country, &take, &after}, locations, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.MeasurementsByCountry(tt.args.ctx, tt.args.country, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.MeasurementsByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.MeasurementsByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_MeasurementsByCity(t *testing.T) {
	take := 1
	after := "mock"
	city := "Paris"
	location := model.LocationResult{ID: "test"}
	locations := []*model.LocationResult{&location}
	locResService.On("GetResultsByCity", city, &take, &after).Return(locations, nil)
	resolver := &Resolver{LocResultsService: locResService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		city  string
		take  *int
		after *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LocationResult
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO(), city, &take, &after}, locations, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.MeasurementsByCity(tt.args.ctx, tt.args.city, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.MeasurementsByCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.MeasurementsByCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_CitiesStartsWith(t *testing.T) {
	city := model.City{ID: "test", Name: "MockCity"}
	cities := []*model.City{&city}
	searchString := "Mock"
	citiesService.On("CitiesStartsWith", searchString).Return(cities, nil)
	resolver := &Resolver{CitiesService: citiesService}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx          context.Context
		searchString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.City
		wantErr bool
	}{
		{"standard", fields{resolver}, args{context.TODO(), searchString}, cities, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CitiesStartsWith(tt.args.ctx, tt.args.searchString)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.CitiesStartsWith() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.CitiesStartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

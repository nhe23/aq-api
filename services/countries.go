package services

import (
	aqdb "github.com/nhe23/aq-api/db"
	"github.com/nhe23/aq-api/graph/model"
)

type CountriesSerivce interface {
	GetCountries() ([]*model.Country, error)
}

type countriesService struct{}

func NewCountriesService() CountriesSerivce {
	return citiesService{}
}

func (s citiesService) GetCountries() ([]*model.Country, error) {
	return aqdb.GetCountries()
}

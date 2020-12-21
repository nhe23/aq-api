package services

import (
	aqdb "github.com/nhe23/aq-api/db"
	"github.com/nhe23/aq-api/graph/model"
)

type CitiesSerivce interface {
	GetCities(take *int, after *string) ([]*model.City, error)
}

type citiesService struct{}

func NewCitiesService() CitiesSerivce {
	return citiesService{}
}

func (s citiesService) GetCities(take *int, after *string) ([]*model.City, error) {
	return aqdb.GetCities(take, after)
}

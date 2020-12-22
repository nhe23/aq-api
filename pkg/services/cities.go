package services

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

type CitiesSerivce interface {
	GetCities(take *int, after *string) ([]*model.City, error)
}

type citiesService struct {
	col dbacc.DataAccess
}

func NewCitiesService(col dbacc.DataAccess) CitiesSerivce {
	return citiesService{col}
}

func (s citiesService) GetCities(take *int, after *string) ([]*model.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := db.GetPaginatedResult(ctx, s.col, take, after)
	if err != nil {
		return nil, err
	}
	var cities []*model.City
	cur.All(ctx, &cities)

	return cities, nil
}

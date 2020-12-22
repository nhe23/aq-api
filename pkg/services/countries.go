package services

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

// CountriesSerivce interface registering all methods
type CountriesSerivce interface {
	GetCountries() ([]*model.Country, error)
}

type countriesService struct {
	col dbacc.DataAccess
}

// NewCountriesService returns CountriesService instance
func NewCountriesService(col dbacc.DataAccess) CountriesSerivce {
	return countriesService{col}
}

func (s countriesService) GetCountries() ([]*model.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := db.GetPaginatedResult(ctx, s.col, nil, nil)
	if err != nil {
		return nil, err
	}
	var countries []*model.Country
	cur.All(ctx, &countries)

	return countries, nil
}

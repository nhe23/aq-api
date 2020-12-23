package countries

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

// Service interface registering all methods
type Service interface {
	GetCountries() ([]*model.Country, error)
}

type service struct {
	col dbacc.DataAccess
}

// NewService returns CountriesService instance
func NewService(col dbacc.DataAccess) Service {
	return service{col}
}

// GetCountries returns paginated country results
func (s service) GetCountries() ([]*model.Country, error) {
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

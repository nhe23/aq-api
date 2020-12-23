package cities

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

// Service interface registering all methods
type Service interface {
	GetCities(take *int, after *string) ([]*model.City, error)
}

type service struct {
	col dbacc.DataAccess
}

// NewService returns new instace of Service
func NewService(col dbacc.DataAccess) Service {
	return service{col}
}

// GetCities returns paginated result of cities
func (s service) GetCities(take *int, after *string) ([]*model.City, error) {
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

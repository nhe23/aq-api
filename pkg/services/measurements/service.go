package measurements

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

// Service interface registering all methods
type Service interface {
	GetResults(take *int, after *string) ([]*model.LocationResult, error)
}

type service struct {
	col dbacc.DataAccess
}

// NewService reruns new LocResService
func NewService(col dbacc.DataAccess) Service {
	return service{col}
}

// GetResults returns paginated location results
func (s service) GetResults(take *int, after *string) ([]*model.LocationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cur, err := db.GetPaginatedResult(ctx, s.col, take, after)
	if err != nil {
		return nil, err
	}
	var resSlice []*model.LocationResult
	cur.All(ctx, &resSlice)

	return resSlice, nil
}

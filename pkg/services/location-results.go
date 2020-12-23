package services

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

// LocResService interface registering all methods
type LocResService interface {
	GetResults(take *int, after *string) ([]*model.LocationResult, error)
}

type locResService struct {
	col dbacc.DataAccess
}

// NewLocResService reruns new LocResService
func NewLocResService(col dbacc.DataAccess) LocResService {
	return locResService{col}
}

// GetResults returns paginated location results
func (s locResService) GetResults(take *int, after *string) ([]*model.LocationResult, error) {
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

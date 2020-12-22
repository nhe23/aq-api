package services

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

type LocResService interface {
	GetResults(take *int, after *string) ([]*model.LocationResult, error)
}

type locResService struct {
	col dbacc.DataAccess
}

func NewLocResService(col dbacc.DataAccess) LocResService {
	return locResService{col}
}

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

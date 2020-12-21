package services

import (
	aqdb "github.com/nhe23/aq-api/db"
	"github.com/nhe23/aq-api/graph/model"
)

type LocResService interface {
	GetResults(take *int, after *string) ([]*model.LocationResult, error)
}

type locResService struct {
}

func NewLocResService() LocResService {
	return locResService{}
}

func (s locResService) GetResults(take *int, after *string) ([]*model.LocationResult, error) {
	return aqdb.LocationResults(take, after)
}

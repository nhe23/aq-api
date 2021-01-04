package cities

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/nhe23/aq-api/graph/model"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetCities(take *int, after *string) (cities []*model.City, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCities",
			"take", take,
			"after", after,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetCities(take, after)
}

func (s *loggingService) CitiesStartsWith(searchString string) (cities []*model.City, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCities",
			"searchstring", searchString,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CitiesStartsWith(searchString)
}

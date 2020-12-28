package measurements

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

func (s *loggingService) GetResults(take *int, after *string) (cities []*model.LocationResult, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetResults",
			"take", take,
			"after", after,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetResults(take, after)
}

func (s *loggingService) GetResultsByCity(city string, take *int, after *string) (cities []*model.LocationResult, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetResultsByCity",
			"take", take,
			"after", after,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetResultsByCity(city, take, after)
}

func (s *loggingService) GetResultsByCountry(country string, take *int, after *string) (cities []*model.LocationResult, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetResultsByCity",
			"take", take,
			"after", after,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetResultsByCountry(country, take, after)
}

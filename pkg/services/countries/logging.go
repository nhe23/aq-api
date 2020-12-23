package countries

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

func (s *loggingService) GetCountries() (cities []*model.Country, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCountries",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetCountries()
}

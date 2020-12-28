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

func (s *loggingService) GetCountries() (countries []*model.Country, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCountries",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetCountries()
}

func (s *loggingService) GetCountry(countryCode string) (country *model.Country, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCountry",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetCountry(countryCode)
}

func (s *loggingService) GetCountryByCodes(countryCodes []string) (country []*model.Country, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetCountryByCodes",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetCountryByCodes(countryCodes)
}

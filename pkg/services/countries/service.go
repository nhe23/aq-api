package countries

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
	"go.mongodb.org/mongo-driver/bson"
)

// Service interface registering all methods
type Service interface {
	GetCountries() ([]*model.Country, error)
	GetCountry(countryCode string) (*model.Country, error)
	GetCountryByCodes(countryCodes []string) ([]*model.Country, error)
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
	filter := bson.M{}
	cur, err := db.GetFilteredResult(ctx, s.col, filter, nil)
	if err != nil {
		return nil, err
	}
	var countries []*model.Country
	cur.All(ctx, &countries)

	return countries, nil
}

// GetCities returns paginated result of cities
func (s service) GetCountry(countryCode string) (*model.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"code": countryCode}

	limit := 1
	cur, err := db.GetFilteredResult(ctx, s.col, filter, &limit)
	if err != nil {
		return nil, err
	}
	var country *model.Country
	cur.Next(ctx)
	cur.Decode(&country)

	return country, nil
}

// GetCities returns paginated result of cities
func (s service) GetCountryByCodes(countryCodes []string) ([]*model.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"code": bson.M{"$in": countryCodes}}
	cur, err := db.GetFilteredResult(ctx, s.col, filter, nil)
	if err != nil {
		return nil, err
	}

	countryByCode := map[string]*model.Country{}
	for cur.Next(ctx) {
		country := model.Country{}
		err := cur.Decode(&country)
		if err != nil {
			return nil, err
		}
		countryByCode[country.Code] = &country
	}
	countries := make([]*model.Country, len(countryCodes))
	for i, id := range countryCodes {
		countries[i] = countryByCode[id]
	}

	return countries, nil
}

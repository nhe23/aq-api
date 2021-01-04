package cities

import (
	"context"
	"fmt"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service interface registering all methods
type Service interface {
	GetCities(take *int, after *string) ([]*model.City, error)
	CitiesStartsWith(searchString string) ([]*model.City, error)
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
	filter := db.GetBasicPaginationFilter(after)
	cur, err := db.GetFilteredResult(ctx, s.col, filter, take)
	if err != nil {
		return nil, err
	}
	var cities []*model.City
	cur.All(ctx, &cities)

	return cities, nil
}

func (s service) CitiesStartsWith(searchString string) ([]*model.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	limit := 20
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{
		Pattern: fmt.Sprintf("^%s", searchString),
		Options: "i",
	}}}
	cur, err := db.GetFilteredResult(ctx, s.col, filter, &limit)
	if err != nil {
		return nil, err
	}
	var cities []*model.City
	cur.All(ctx, &cities)

	return cities, nil
}

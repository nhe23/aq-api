package measurements

import (
	"context"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/db"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service interface registering all methods
type Service interface {
	GetResults(take *int, after *string) ([]*model.LocationResult, error)
	GetResultsByCity(city string, take *int, after *string) ([]*model.LocationResult, error)
	GetResultsByCountry(country string, take *int, after *string) ([]*model.LocationResult, error)
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
	filter := db.GetBasicPaginationFilter(after)

	cur, err := db.GetFilteredResult(ctx, s.col, filter, take)
	if err != nil {
		return nil, err
	}
	var resSlice []*model.LocationResult
	cur.All(ctx, &resSlice)
	return resSlice, nil
}

// GetResults returns paginated location results
func (s service) GetResultsByCountry(country string, take *int, after *string) ([]*model.LocationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var filter primitive.M
	if after != nil {
		docID, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": docID}, "country": country}
	} else {
		filter = bson.M{"country": country}
	}

	cur, err := db.GetFilteredResult(ctx, s.col, filter, take)
	if err != nil {
		return nil, err
	}
	var resSlice []*model.LocationResult
	cur.All(ctx, &resSlice)
	return resSlice, nil
}

// GetResults returns paginated location results
func (s service) GetResultsByCity(city string, take *int, after *string) ([]*model.LocationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var filter primitive.M
	if after != nil {
		docID, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": docID}, "city": city}
	} else {
		filter = bson.M{"city": city}
	}

	cur, err := db.GetFilteredResult(ctx, s.col, filter, take)
	if err != nil {
		return nil, err
	}
	var resSlice []*model.LocationResult
	cur.All(ctx, &resSlice)
	return resSlice, nil
}

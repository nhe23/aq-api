package db

import (
	"context"

	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type DataAccessCursor interface {
// 	All(ctx, results interface{}) error
// }

// // DataAccess interface is used to abstract db dependency
// type DataAccess interface {
// 	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (DataAccessCursor, error)
// }

// GetPaginatedResult return db cursor with paginated results
func GetPaginatedResult(ctx context.Context, col dbacc.DataAccess, take *int, after *string) (*mongo.Cursor, error) {
	options := options.Find()
	if take != nil {
		options.SetLimit(int64(*take))
	}

	var filter primitive.M
	if after != nil {
		docID, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": docID}}
	} else {
		filter = bson.M{}
	}
	cur, err := col.Find(ctx, filter, options)
	return cur, err
}

// GetFilteredResult returns collection cursor for given filter
func GetFilteredResult(ctx context.Context, col dbacc.DataAccess, filter primitive.M, limit *int) (*mongo.Cursor, error) {
	options := options.Find()
	if limit != nil {
		options.SetLimit(int64(*limit))
	}
	cur, err := col.Find(ctx, filter, options)
	return cur, err
}

func GetBasicPaginationFilter(after *string) primitive.M {
	var filter primitive.M
	if after != nil {
		docID, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": docID}}
	} else {
		filter = bson.M{}
	}
	return filter
}

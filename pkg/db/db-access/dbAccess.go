package dbacc

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type DataAccessCursor interface {
// 	Decode(val interface{}) error
// 	Err() error
// 	Next(ctx context.Context) bool
// 	Close(ctx context.Context) error
// 	ID() int64
// 	Current() bson.Raw
// 	All(ctx context.Context, results interface{}) error
// }

type mongoCursor struct {
	mongo.Cursor
}

func (m *mongoCursor) Current() bson.Raw {
	return m.Cursor.Current
}

// DataAccess interface is used to abstract db dependency
type DataAccess interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

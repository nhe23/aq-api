package mock

import (
	"context"
	"reflect"

	"bou.ke/monkey"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCursor struct{}

// type DataAccessCursor interface {
// 	All(ctx, results interface{}) error
// }

// func NewMockCursor() dbacc.DataAccessCursor {
// 	return mockCursor{}
// }

func (m mockCursor) All(ctx, results interface{}) error {
	return nil
}

type mockDataAccess struct{}

// DataAccess interface is used to abstract db dependency
type DataAccess interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

// NewMockDataAccess returns mocked db data access
func NewMockDataAccess() DataAccess {
	var c *mongo.Cursor
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(c), "All",
		func(c *mongo.Cursor, ctx context.Context, results interface{}) error {
			guard.Unpatch()
			return nil
		})
	return mockDataAccess{}
}

func (m mockDataAccess) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	// cur := NewMockCursor()
	cur := &mongo.Cursor{}
	return cur, nil
}

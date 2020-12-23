package db

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"

	"github.com/nhe23/aq-api/mocks"
)

func TestGetPaginatedResult(t *testing.T) {
	dataAccess := mocks.NewMockDataAccess()
	ctx := context.TODO()
	nilCursor := &mongo.Cursor{}
	take := 5
	after := "mock"
	type args struct {
		ctx   context.Context
		col   dbacc.DataAccess
		take  *int
		after *string
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Cursor
		wantErr bool
	}{
		{"standard", args{ctx, dataAccess, &take, &after}, nilCursor, false},
		{"nilValues", args{ctx, dataAccess, nil, nil}, nilCursor, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPaginatedResult(tt.args.ctx, tt.args.col, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaginatedResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPaginatedResult() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

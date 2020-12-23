package services

import (
	"reflect"
	"testing"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/mocks"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

func Test_locResService_GetResults(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	take := 5
	after := "mock"
	type fields struct {
		col dbacc.DataAccess
	}
	type args struct {
		take  *int
		after *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LocationResult
		wantErr bool
	}{
		{"standard", fields{mockCol}, args{&take, &after}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := locResService{
				col: tt.fields.col,
			}
			got, err := s.GetResults(tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("locResService.GetResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("locResService.GetResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"reflect"
	"testing"

	"github.com/nhe23/aq-api/graph/model"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"

	"github.com/nhe23/aq-api/mock"
)

func Test_citiesService_GetCities(t *testing.T) {
	mockCol := mock.NewMockDataAccess()
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
		want    []*model.City
		wantErr bool
	}{
		{"standard", fields{mockCol}, args{&take, &after}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := citiesService{
				col: tt.fields.col,
			}
			got, err := s.GetCities(tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("citiesService.GetCities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("citiesService.GetCities() = %v, want %v", got, tt.want)
			}
		})
	}
}

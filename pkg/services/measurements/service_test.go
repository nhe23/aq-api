package measurements

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
			s := service{
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

func Test_service_GetResultsByCountry(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	take := 5
	after := "mock"
	type fields struct {
		col dbacc.DataAccess
	}
	type args struct {
		country string
		take    *int
		after   *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.LocationResult
		wantErr bool
	}{
		{"standard", fields{mockCol}, args{"DE", &take, &after}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				col: tt.fields.col,
			}
			got, err := s.GetResultsByCountry(tt.args.country, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetResultsByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetResultsByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetResultsByCity(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	take := 5
	after := "mock"
	type fields struct {
		col dbacc.DataAccess
	}
	type args struct {
		city  string
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
		{"standard", fields{mockCol}, args{"Paris", &take, &after}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				col: tt.fields.col,
			}
			got, err := s.GetResultsByCity(tt.args.city, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetResultsByCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetResultsByCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

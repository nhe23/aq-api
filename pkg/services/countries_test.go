package services

import (
	"reflect"
	"testing"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/mocks"
	dbacc "github.com/nhe23/aq-api/pkg/db/db-access"
)

func Test_countriesService_GetCountries(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	type fields struct {
		col dbacc.DataAccess
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.Country
		wantErr bool
	}{
		{"standard", fields{mockCol}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := countriesService{
				col: tt.fields.col,
			}
			got, err := s.GetCountries()
			if (err != nil) != tt.wantErr {
				t.Errorf("countriesService.GetCountries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countriesService.GetCountries() = %v, want %v", got, tt.want)
			}
		})
	}
}

package countries

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
			s := service{
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

func Test_service_GetCountry(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	type fields struct {
		col dbacc.DataAccess
	}
	type args struct {
		countryCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Country
		wantErr bool
	}{
		{"standard", fields{mockCol}, args{"asdf"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				col: tt.fields.col,
			}
			got, err := s.GetCountry(tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetCountryByCodes(t *testing.T) {
	mockCol := mocks.NewMockDataAccess()
	countryCodes := []string{"DE", "FR", "US"}
	want := []*model.Country{nil, nil, nil}
	type fields struct {
		col dbacc.DataAccess
	}
	type args struct {
		countryCodes []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Country
		wantErr bool
	}{
		{"standard", fields{mockCol}, args{countryCodes}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				col: tt.fields.col,
			}
			got, err := s.GetCountryByCodes(tt.args.countryCodes)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetCountryByCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetCountryByCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

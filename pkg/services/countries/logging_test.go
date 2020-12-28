package countries

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/mocks"
)

var take int = 1
var after string = "mock"

func Test_loggingService_GetCountries(t *testing.T) {
	country := model.Country{ID: "test"}
	countries := []*model.Country{&country}
	service := new(mocks.CountriesService)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	service.On("GetCountries").Return(countries, nil)
	type fields struct {
		logger  log.Logger
		Service Service
	}
	tests := []struct {
		name          string
		fields        fields
		wantCountries []*model.Country
		wantErr       bool
	}{
		{"standard", fields{logger, service}, countries, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCountries, err := s.GetCountries()
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetCountries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountries, tt.wantCountries) {
				t.Errorf("loggingService.GetCountries() = %v, want %v", gotCountries, tt.wantCountries)
			}
		})
	}
}

func Test_loggingService_GetCountry(t *testing.T) {
	country := model.Country{ID: "test"}
	service := new(mocks.CountriesService)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	countryParam := "DE"
	service.On("GetCountry", countryParam).Return(&country, nil)
	type fields struct {
		logger  log.Logger
		Service Service
	}
	type args struct {
		countryCode string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCountry *model.Country
		wantErr     bool
	}{
		{"standard", fields{logger, service}, args{countryParam}, &country, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCountry, err := s.GetCountry(tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountry, tt.wantCountry) {
				t.Errorf("loggingService.GetCountry() = %v, want %v", gotCountry, tt.wantCountry)
			}
		})
	}
}

func Test_loggingService_GetCountryByCodes(t *testing.T) {
	countryCodes := []string{"DE", "FR"}
	country := model.Country{ID: "test"}
	countries := []*model.Country{&country}
	service := new(mocks.CountriesService)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	service.On("GetCountryByCodes", countryCodes).Return(countries, nil)
	type fields struct {
		logger  log.Logger
		Service Service
	}
	type args struct {
		countryCodes []string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCountry []*model.Country
		wantErr     bool
	}{
		{"standard", fields{logger, service}, args{countryCodes}, countries, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCountry, err := s.GetCountryByCodes(tt.args.countryCodes)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetCountryByCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountry, tt.wantCountry) {
				t.Errorf("loggingService.GetCountryByCodes() = %v, want %v", gotCountry, tt.wantCountry)
			}
		})
	}
}

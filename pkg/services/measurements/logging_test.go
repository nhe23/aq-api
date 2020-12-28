package measurements

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

func Test_loggingService_GetResults(t *testing.T) {
	result := model.LocationResult{ID: "test"}
	results := []*model.LocationResult{&result}
	service := new(mocks.MeasurementsService)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	service.On("GetResults", &take, &after).Return(results, nil)
	after := "mock"
	type fields struct {
		logger  log.Logger
		Service Service
	}
	type args struct {
		take  *int
		after *string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCities []*model.LocationResult
		wantErr    bool
	}{
		{"standard", fields{logger, service}, args{&take, &after}, results, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCities, err := s.GetResults(tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCities, tt.wantCities) {
				t.Errorf("loggingService.GetResults() = %v, want %v", gotCities, tt.wantCities)
			}
		})
	}
}

func Test_loggingService_GetResultsByCity(t *testing.T) {
	result := model.LocationResult{ID: "test"}
	results := []*model.LocationResult{&result}
	service := new(mocks.MeasurementsService)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	city := "Paris"
	service.On("GetResultsByCity", city, &take, &after).Return(results, nil)
	after := "mock"
	type fields struct {
		logger  log.Logger
		Service Service
	}
	type args struct {
		city  string
		take  *int
		after *string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCities []*model.LocationResult
		wantErr    bool
	}{
		{"standard", fields{logger, service}, args{city, &take, &after}, results, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCities, err := s.GetResultsByCity(tt.args.city, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetResultsByCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCities, tt.wantCities) {
				t.Errorf("loggingService.GetResultsByCity() = %v, want %v", gotCities, tt.wantCities)
			}
		})
	}
}

func Test_loggingService_GetResultsByCountry(t *testing.T) {
	result := model.LocationResult{ID: "test"}
	results := []*model.LocationResult{&result}
	service := new(mocks.MeasurementsService)
	country := "DE"
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	service.On("GetResultsByCountry", country, &take, &after).Return(results, nil)
	after := "mock"
	type fields struct {
		logger  log.Logger
		Service Service
	}
	type args struct {
		country string
		take    *int
		after   *string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCities []*model.LocationResult
		wantErr    bool
	}{
		{"standard", fields{logger, service}, args{country, &take, &after}, results, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCities, err := s.GetResultsByCountry(tt.args.country, tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetResultsByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCities, tt.wantCities) {
				t.Errorf("loggingService.GetResultsByCountry() = %v, want %v", gotCities, tt.wantCities)
			}
		})
	}
}

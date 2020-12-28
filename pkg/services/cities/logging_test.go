package cities

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

func Test_loggingService_GetCities(t *testing.T) {
	city := model.City{ID: "test", Name: "MockCity"}
	cities := []*model.City{&city}
	service := new(mocks.CitiesSerivce)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	service.On("GetCities", &take, &after).Return(cities, nil)
	// mockCol := mocks.NewMockDataAccess()
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
		wantCities []*model.City
		wantErr    bool
	}{
		{"standard", fields{logger, service}, args{&take, &after}, cities, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &loggingService{
				logger:  tt.fields.logger,
				Service: tt.fields.Service,
			}
			gotCities, err := s.GetCities(tt.args.take, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingService.GetCities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCities, tt.wantCities) {
				t.Errorf("loggingService.GetCities() = %v, want %v", gotCities, tt.wantCities)
			}
		})
	}
}

package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"github.com/nhe23/aq-api/pkg/services/countries"
)

//go:generate go run github.com/vektah/dataloaden CountryLoader string *github.com/nhe23/aq-api/graph/model.Country

const loadersKey = "dataloaders"

type Loader interface {
	For(ctx context.Context) *Loaders
}

type loader struct{}

// NewLoader returns new instace of Loader
func NewLoader() Loader {
	return loader{}
}

// Loaders type
type Loaders struct {
	CountryByCode CountryLoader
}

//Middleware returns a handler that registers different loaders
func Middleware(countriesService countries.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			CountryByCode: CountryLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(countryCodes []string) ([]*model.Country, []error) {
					errors := make([]error, 0)
					countries, err := countriesService.GetCountryByCodes(countryCodes)
					errors = append(errors, err)
					return countries, errors
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//For returns registered Loaders for given key
func (l loader) For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

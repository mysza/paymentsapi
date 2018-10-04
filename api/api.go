package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API provides the application HTTP API
type API struct {
	// API resources needed to work
	Payments *PaymentResource
}

// NewAPI creates a new API instance
func NewAPI( /* inject deps */ ) (*API, error) {
	payments := NewPaymentResource(nil)
	return &API{
		Payments: payments,
	}, nil
}

// Router provides API routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(middleware.DefaultCompress)
	r.Mount("/payments", api.Payments.router())
	return r
}

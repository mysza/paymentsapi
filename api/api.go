package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// API provides the application HTTP API
type API struct {
	payments *PaymentResource
	router   *chi.Mux
}

// NewAPI creates a new API instance
func NewAPI() (*API, error) {
	payments := NewPaymentResource(nil)
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.DefaultCompress)
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(middleware.DefaultCompress)
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Mount("/payments", payments.router())
	return &API{payments, router}, nil
}

// Router provides API routes.
func (api *API) Router() *chi.Mux {
	return api.router
}

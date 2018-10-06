package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Server is the HTTP server of the Payments API.
type Server struct {
	*http.Server
}

// NewServer creates and configures a new API server for all application routes.
func NewServer(address string) (*Server, error) {
	api, err := NewAPI()
	if err != nil {
		return nil, err
	}
	srv := http.Server{Addr: address, Handler: api.Router()}
	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}

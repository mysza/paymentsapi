package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger"
	"github.com/mysza/paymentsapi/repository"
)

// Server is the HTTP server of the Payments API.
type Server struct {
	*http.Server
}

func createDatabase() (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = "./db"
	opts.ValueDir = "./db"
	return badger.Open(opts)
}

// StartHTTPServer starts HTTP server on a given port, with database
// being used at dbDir.
func StartHTTPServer(port, dbDir string) error {
	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.New(db)
	api, err := NewAPI(repo)
	if err != nil {
		log.Fatal(err)
	}
	srv := http.Server{Addr: fmt.Sprintf("localhost:%v", port), Handler: api.Router()}
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
	log.Println("Server stopped")
	return nil
}

package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger"
	"github.com/mysza/paymentsapi/repository"
	"github.com/sirupsen/logrus"
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
		logrus.WithFields(logrus.Fields{
			"location": "api/server/StartHTTPServer",
			"details":  "database open",
			"error":    err,
		}).Error("Error opening database")
		panic(err)
	}
	defer db.Close()
	repo := repository.New(db)
	api, err := NewAPI(repo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "api/server/StartHTTPServer",
			"details":  "api creation",
			"error":    err,
		}).Error("Error creating API")
		panic(err)
	}
	srv := http.Server{Addr: fmt.Sprintf("localhost:%v", port), Handler: api.Router()}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithFields(logrus.Fields{
				"location": "api/server/StartHTTPServer",
				"details":  "listen & serve",
				"error":    err,
			}).Error("Error starting HTTP Server")
			panic(err)
		}
	}()
	logrus.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logrus.Println("Shutting down server... Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	logrus.Println("Server stopped")
	return nil
}

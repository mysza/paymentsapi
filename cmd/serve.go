package cmd

import (
	"log"

	"github.com/dgraph-io/badger"
	"github.com/mysza/paymentsapi/api"
	"github.com/mysza/paymentsapi/repository"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// blank import to initialize the SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func createDatabase() (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = "./db"
	opts.ValueDir = "./db"
	return badger.Open(opts)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := createDatabase()
		if err != nil {
			log.Fatal(err)
		}
		repo := repository.New(db)
		server, err := api.NewServer(viper.GetString("address"), repo)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	viper.SetDefault("address", "localhost:3000")
}

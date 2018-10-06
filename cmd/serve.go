package cmd

import (
	"log"

	"github.com/mysza/paymentsapi/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	// blank import to initialize the SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open("sqlite3", "./db/gorm.db")
		if err != nil {
			log.Fatal(err)
		}
		server, err := api.NewServer(viper.GetString("address"), db)
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

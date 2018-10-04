package cmd

import (
	"log"

	"github.com/mysza/paymentsapi/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.
	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("log_level", "debug")
}

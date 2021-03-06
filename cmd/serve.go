package cmd

import (
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
		api.StartHTTPServer(viper.GetString("port"), viper.GetString("dbdir"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	viper.SetDefault("port", "3000")
	viper.SetDefault("dbdir", "./db")
}

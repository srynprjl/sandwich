package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/srynprjl/sandwich/api"
)

var port int
var host string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start a web server for managing your projects",
	Run: func(cmd *cobra.Command, args []string) {
		if port < 1024 || port > 65535 {
			fmt.Fprintf(os.Stderr, "Error: port must be between 1024 and 65535\n")
			os.Exit(1)
		}
		api.Api("localhost", 5000)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run server on")
	webCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to run server on")
}

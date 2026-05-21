package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yashodhanketkar/arsg/src/api"
)

var port string

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start in RESTful API mode",
	Long:  `Start the application in RESTFul API.`,
	Run:   func(cmd *cobra.Command, args []string) { api.Serve(DB, ":"+port) },
}

func init() {
	rootCmd.AddCommand(restCmd)
	restCmd.Flags().StringVarP(&port, "port", "p", "5000", "Port to run the RESTful API on")
}

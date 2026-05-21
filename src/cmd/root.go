package cmd

import (
	"database/sql"
	"os"

	"github.com/spf13/cobra"
	"github.com/yashodhanketkar/arsg/src/db"
)

var DB *sql.DB
var mode string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "arsg",
	Short: "A small content rating application",
	Long: `A small local binary for rating content.

This application is built using SQLite, BubbleTea, and Chi router libraries.
There are two modes available for running the application:
	- REST: Starts the application with a chi powered RESTful API
	- UI: Starts the application with a BubbleTea powerd TUI interface
`,
	PersistentPreRun:  func(cmd *cobra.Command, args []string) { DB = db.ConnectDB(mode) },
	PersistentPostRun: func(cmd *cobra.Command, args []string) { DB.Close() },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().
		StringVarP(&mode, "mode", "m", "prod", "Mode to run the application in (Options: dev or prod)")

	rootCmd.RegisterFlagCompletionFunc(
		"mode",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"dev", "prod"}, cobra.ShellCompDirectiveNoFileComp
		},
	)
}

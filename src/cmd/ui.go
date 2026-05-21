package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yashodhanketkar/arsg/src/ui"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start in TUI mode",
	Long:  "Start the application in stdio mode with bubbletea UI library",
	Run:   func(cmd *cobra.Command, args []string) { ui.TeaUI(DB) },
}

func init() {
	rootCmd.AddCommand(uiCmd)
}

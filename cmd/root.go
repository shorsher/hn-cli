package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd launches our application
var RootCmd = &cobra.Command{
	Use: "hn-cli",
	Short: "an interactive HN cli",
	Long: "A CLI client to browse HN discreetly from your terminal :)",
}

func init() {
	AddCommands()
}

// AddCommands to your root command
func AddCommands() {
 RootCmd.AddCommand(topCmd)	
}
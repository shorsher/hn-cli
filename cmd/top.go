package cmd

import (
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use: "top",
	Short: "show top HN stories",
	Long: "show the top 30 HN stories, similar to visiting the front page",
}

func init() {
	topCmd.Run = top
}

func top(cmd *cobra.Command, args []string) {
	Top()
}
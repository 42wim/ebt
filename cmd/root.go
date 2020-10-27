package cmd

import (
	"github.com/spf13/cobra"
)

func (c *cmd) rootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "ebt",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	rootCmd.AddCommand(combineCmd())

	return rootCmd
}

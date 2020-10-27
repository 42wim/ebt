package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type cmd struct {
	root *cobra.Command
}

func Execute() {
	c := &cmd{}

	c.root = c.rootCmd()
	if err := c.root.Execute(); err != nil {
		log.Fatalln(err)
	}
}

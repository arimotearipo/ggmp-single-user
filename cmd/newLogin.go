package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:     "new",
    Short:   "Add a new set of login details",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]

        fmt.Printf("Setting up a new set of login details for domain %s\n", domain)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
    Use:     "edit",
    Short:   "Edit existing login details",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        domain := args[0]

        fmt.Printf("Editing existing set of login details for %s\n", domain)
    },
}

func init() {
    rootCmd.AddCommand(editCmd)
}
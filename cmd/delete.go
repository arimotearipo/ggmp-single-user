package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete an existing set of login details",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]

		fmt.Printf("Deleting login details at index %s\n", index)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
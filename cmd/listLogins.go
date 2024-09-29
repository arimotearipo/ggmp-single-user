package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List existing logins",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Println("Listing logins")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
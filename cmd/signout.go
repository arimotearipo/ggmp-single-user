package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)



var signOutCmd = &cobra.Command{
	Use: "signout",
	Short: "Signout of the application",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Println("Signing out of the application")
	},
}

func init() {
	rootCmd.AddCommand(signOutCmd)
}
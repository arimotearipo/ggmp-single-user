package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signInCmd = &cobra.Command{
	Use: "signin",
	Short: "Signin to the application",
	Args: cobra.ExactArgs(2),
	Run: func (cmd *cobra.Command, args []string)  {
		username, password := args[0], args[1]	

		fmt.Printf("Signing in with username %s and password %s\n", username, password)
	},
}

var signOutCmd = &cobra.Command{
	Use: "signout",
	Short: "Signout of the application",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Println("Signing out of the application")
	},
}

func init() {
	rootCmd.AddCommand(signInCmd)
	rootCmd.AddCommand(signOutCmd)
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signInCmd = &cobra.Command{
	Use: "signin",
	Short: "Signin to the application",
	Args: cobra.ExactArgs(0),
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Printf("Signing in with username %s and password %s\n", fUsername, fPassword)
	},
}

func init() {
	signInCmd.Flags().StringVarP(&fUsername, "username", "u", "", "Your username")
	signInCmd.Flags().StringVarP(&fPassword, "password", "p", "", "Your master password")

	rootCmd.AddCommand(signInCmd)
}
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

		if err != nil {
			log.Fatal("Fail to generate password hash")
		}

		q.RegisterAccount(username, string(bytes))
	},
}

func init() {
	registerCmd.Flags().StringVarP(&username, "username", "u", "", "Your username")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "Your master password")

	rootCmd.AddCommand(registerCmd)
}

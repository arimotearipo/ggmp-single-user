package cmd

import (
	"log"

	"github.com/arimotearipo/ggmp/encryption"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var signInCmd = &cobra.Command{
	Use:   "signin",
	Short: "Signin to the application",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		hashedPassword := q.SignInAccount(username)

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			log.Fatal("Incorrect password")
		}

		c = encryption.NewCrypt([]byte(password))
		log.Println("Successfully logged in")
	},
}

func init() {
	signInCmd.Flags().StringVarP(&username, "username", "u", "", "Your username")
	signInCmd.Flags().StringVarP(&password, "password", "p", "", "Your master password")

	rootCmd.AddCommand(signInCmd)
}

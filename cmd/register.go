package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var fUsername string
var fPassword string

var registerCmd = &cobra.Command{
	Use: "register",
	Short: "Create a new account",
	Run: func (cmd *cobra.Command, args []string)  {
		hash := sha256.New()
		hash.Write([]byte(fPassword))

		hashedPassword := hex.EncodeToString(hash.Sum(nil))

		fmt.Println(hashedPassword)
	},
}

func init() {
	registerCmd.Flags().StringVarP(&fUsername, "username", "u", "", "Your username")
	registerCmd.Flags().StringVarP(&fPassword, "password", "p", "", "Your master password")


	rootCmd.AddCommand(registerCmd)
}
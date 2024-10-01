package cmd

import (
	"fmt"
	"os"

	"github.com/arimotearipo/ggmp/encryption"
	"github.com/arimotearipo/ggmp/queries"
	"github.com/spf13/cobra"
)

var c *encryption.Cryption
var q *queries.Queries
var username string
var password string

func SetQueries(_q *queries.Queries) {
	q = _q
}

var rootCmd = &cobra.Command{
	Use:   "ggmp",
	Short: "ggmp - go get my password",
	Long:  `A CLI application to store and manage your passwords -- built with Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running ggmp")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/arimotearipo/ggmp/cmd"
	"github.com/arimotearipo/ggmp/database"
)

func readCommands(c *cmd.Command) {
	for {
		fmt.Println("1. Get password")
		fmt.Println("2. Add password")
		fmt.Println("3. List URIs")
		fmt.Println("4. Delete password")
		fmt.Println("5. Update password")
		fmt.Println("6. Exit")
		fmt.Printf("Enter your choice: ")

		var choice int
		fmt.Scanf("%d", &choice)
		fmt.Println(choice)

		switch choice {
		case 1:
			c.GetPassword()
		case 2:
			c.AddPassword()
		case 3:
			c.ListURIs()
		case 4:
			c.DeletePassword()
		case 5:
			c.UpdatePassword()
		case 6:
			fmt.Println("Exit")
			os.Exit(0)
		}
	}
}

func authenticate(c *cmd.Command) {
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. List accounts")
		fmt.Println("4. Delete account")
		fmt.Println("5. Exit")
		fmt.Printf("Enter your choice: ")

		var choice int
		fmt.Scanf("%d", &choice)
		fmt.Println(choice)

		switch choice {
		case 1:
			if c.Login() {
				return
			}
		case 2:
			if c.Register() {
				continue
			}
		case 3:
			c.ListAccounts()
		case 4:
			c.Delete()
		case 5:
			fmt.Println("Exit")
			os.Exit(0)
		}
	}
}

func main() {
	fmt.Println("Welcome to GGMP CLI")

	db := database.NewDatabase("ggmp.db")
	defer db.Close()

	commands := cmd.NewCommands(db)

	authenticate(commands)

	readCommands(commands)
}

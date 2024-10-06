package main

import (
	"fmt"
	"ggmp/cmd"
	"ggmp/database"
	"ggmp/encryption"
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
			return
		}
	}
}

func main() {
	fmt.Println("Welcome to GGMP CLI")

	e := encryption.NewEncryption([]byte("password"))

	db := database.NewDatabase("ggmp.db")
	defer db.Close()

	commands := cmd.NewCommands(e, db)

	readCommands(commands)
}

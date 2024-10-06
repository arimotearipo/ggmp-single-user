package main

import (
	"fmt"
	"ggmp/database"
	"ggmp/encryption"
)

var e *encryption.Encryption
var db *database.Database

func addPassword() {
	fmt.Println("Adding password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	fmt.Println("Enter username:")
	var username string
	fmt.Scanf("%s", &username)

	fmt.Println("Enter password:")
	var password string
	fmt.Scanf("%s", &password)

	encryptedPassword, err := e.Encrypt(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AddPassword(uri, username, encryptedPassword)
}

func getPassword() {
	fmt.Println("Getting password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	encryptedPassword, err := db.GetPassword(uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	password, err := e.Decrypt(encryptedPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Password: ", password)
}

func listURIs() {
	fmt.Println("Listing URIs")

	uris, err := db.ListURIs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("URIs: ", uris)
}

func deletePassword() {
	fmt.Println("Deleting password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	db.DeleteAccount(uri)
}

func updatePassword() {
	fmt.Println("Updating password")

	fmt.Printf("Enter uri: ")
	var uri string
	fmt.Scanf("%s", &uri)

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanf("%s", &username)

	fmt.Printf("Enter password: ")
	var password string
	fmt.Scanf("%s", &password)

	hashed_password, err := e.Encrypt(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.UpdatePassword(uri, username, hashed_password)
}

func readCommands() {
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
			getPassword()
		case 2:
			addPassword()
		case 3:
			listURIs()
		case 4:
			deletePassword()
		case 5:
			updatePassword()
		case 6:
			fmt.Println("Exit")
			return
		}
	}
}

func main() {
	fmt.Println("Welcome to GGMP CLI")

	e = encryption.NewEncryption([]byte("password"))

	db = database.NewDatabase("ggmp.db")

	readCommands()
}

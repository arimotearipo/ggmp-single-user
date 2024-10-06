package cmd

import (
	"fmt"
	"ggmp/database"
	"ggmp/encryption"
)

type Command struct {
	e  *encryption.Encryption
	db *database.Database
}

func NewCommands(e *encryption.Encryption, db *database.Database) *Command {
	return &Command{e: e, db: db}
}

func (c *Command) AddPassword() {
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

	encryptedPassword, err := c.e.Encrypt(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.db.AddPassword(uri, username, encryptedPassword)
}

func (c *Command) GetPassword() {
	fmt.Println("Getting password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	encryptedPassword, err := c.db.GetPassword(uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	password, err := c.e.Decrypt(encryptedPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Password: ", password)
}

func (c *Command) ListURIs() {
	fmt.Println("Listing URIs")

	uris, err := c.db.ListURIs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("URIs: ", uris)
}

func (c *Command) DeletePassword() {
	fmt.Println("Deleting password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	c.db.DeleteAccount(uri)
}

func (c *Command) UpdatePassword() {
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

	hashed_password, err := c.e.Encrypt(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.db.UpdatePassword(uri, username, hashed_password)
}

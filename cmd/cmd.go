package cmd

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"ggmp/database"
	"ggmp/encryption"

	"golang.org/x/crypto/bcrypt"
)

type Command struct {
	e  *encryption.Encryption
	db *database.Database
}

func NewCommands(db *database.Database) *Command {
	return &Command{e: nil, db: db}
}

// === LOGINS ====
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

// === MASTER ACCOUNT ===
func (c *Command) Login() bool {
	fmt.Println("Login")

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanf("%s", &username)

	fmt.Printf("Enter password: ")
	var password string
	fmt.Scanf("%s", &password)

	hashedPassword, initializationVector, salt, err := c.db.GetMasterAccount(username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}

	c.e = encryption.NewEncryption([]byte(password), initializationVector, salt)
	return true
}

func (c *Command) Register() bool {
	fmt.Println("Register")

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanf("%s", &username)

	fmt.Printf("Enter psasword: ")
	var password string
	fmt.Scanf("%s", &password)

	initializationVector := make([]byte, aes.BlockSize)
	_, err := rand.Read(initializationVector)
	if err != nil {
		fmt.Println(err)
		return false
	}

	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		fmt.Println(err)
		return false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println(err)
		return false
	}

	c.db.AddMasterAccount(username, string(hashedPassword), initializationVector, salt)

	fmt.Println("Signup successful. Proceed to login.")
	return true
}

func (c *Command) Delete() bool {
	fmt.Println("Delete")

	fmt.Printf("Enter username:")
	var username string
	fmt.Scanf("%s", &username)

	err := c.db.DeleteMasterAccount(username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Account deleted successfully.")
	return true
}

func (c *Command) ListAccounts() {
	fmt.Println("List accounts")

	accounts, err := c.db.ListMasterAccounts()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Accounts: ", accounts)
}

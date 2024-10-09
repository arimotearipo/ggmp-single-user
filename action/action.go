package action

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"

	"github.com/arimotearipo/ggmp/database"
	"github.com/arimotearipo/ggmp/encryption"

	"golang.org/x/crypto/bcrypt"
)

type Action struct {
	e  *encryption.Encryption
	db *database.Database
}

func NewAction(db *database.Database) *Action {
	return &Action{e: nil, db: db}
}

// === LOGINS ====
func (c *Action) AddPassword() {
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

func (c *Action) GetPassword() {
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

func (c *Action) ListURIs() {
	fmt.Println("Listing URIs")

	uris, err := c.db.ListURIs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("URIs: ", uris)
}

func (c *Action) DeletePassword() {
	fmt.Println("Deleting password")

	fmt.Println("Enter uri:")
	var uri string
	fmt.Scanf("%s", &uri)

	c.db.DeleteAccount(uri)
}

func (c *Action) UpdatePassword() {
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

// Will prompt the user for username and password and
// proceeds to compare the hash and password
func (c *Action) Login(username, password string) (bool, error) {
	hashedPassword, initializationVector, salt, err := c.db.GetMasterAccount(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	c.e = encryption.NewEncryption([]byte(password), initializationVector, salt)
	return true, nil
}

func (c *Action) Register(username, password string) string {
	initializationVector := make([]byte, aes.BlockSize)
	_, err := rand.Read(initializationVector)
	if err != nil {
		return err.Error()
	}

	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return err.Error()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err.Error()
	}

	err = c.db.AddMasterAccount(username, string(hashedPassword), initializationVector, salt)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (c *Action) Delete(username, password string) (bool, error) {
	if ok, err := c.Login(username, password); !ok || err != nil {
		return false, err
	}

	err := c.db.DeleteMasterAccount(username)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Action) ListAccounts() []string {
	accounts, err := c.db.ListMasterAccounts()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return accounts
}

package action

import (
	"crypto/rand"

	"github.com/arimotearipo/ggmp/internal/database"
	"github.com/arimotearipo/ggmp/internal/encryption"

	"golang.org/x/crypto/bcrypt"
)

type session struct {
	username  string
	masterKey []byte
	id        int
}

type Action struct {
	db   *database.Database
	sess *session
}

func NewAction(db *database.Database) *Action {
	return &Action{db: db, sess: nil}
}

// === LOGINS ====
func (a *Action) AddPassword(uri, username, password string) error {
	encryptedPassword, err := encryption.Encrypt(password, a.sess.masterKey)
	if err != nil {
		return err
	}

	err = a.db.AddPassword(a.sess.id, uri, username, encryptedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) GetPassword(uri string) (string, string, error) {
	username, encryptedPassword, err := a.db.GetPassword(uri, a.sess.id)
	if err != nil {
		return "", "", err
	}

	password, err := encryption.Decrypt(encryptedPassword, a.sess.masterKey)
	if err != nil {
		return "", "", err
	}

	return username, password, nil
}

func (a *Action) ListURIs() ([]string, error) {
	uris, err := a.db.ListURIs(a.sess.id)
	if err != nil {
		return nil, err
	}

	return uris, nil
}

func (a *Action) DeletePassword(uri string) error {
	err := a.db.DeleteAccount(uri, a.sess.id)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) UpdatePassword(uri, username, password string) error {
	hashed_password, err := encryption.Encrypt(password, a.sess.masterKey)
	if err != nil {
		return err
	}

	err = a.db.UpdatePassword(uri, username, hashed_password)
	if err != nil {
		return err
	}

	return nil
}

// === MASTER ACCOUNT ===

// Will prompt the user for username and password and
// proceeds to compare the hash and password
func (a *Action) Login(username, password string) error {
	hashedPassword, salt, err := a.db.GetMasterAccount(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	id, err := a.db.GetUserId(username)
	if err != nil {
		return err
	}

	masterKey := encryption.DeriveKey([]byte(password), salt)

	a.sess = &session{username, masterKey, id}

	return nil
}

func (a *Action) Logout() {
	a.sess.username = ""
	a.sess.masterKey = nil
	a.sess.id = -1
	a.sess = nil
}

func (a *Action) Register(username, password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err.Error()
	}

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err.Error()
	}

	err = a.db.AddMasterAccount(username, string(hashedPassword), salt)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (a *Action) Delete(username, password string) error {
	if err := a.Login(username, password); err != nil {
		return err
	}

	err := a.db.DeleteMasterAccount(username)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) ListAccounts() ([]string, error) {
	accounts, err := a.db.ListMasterAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

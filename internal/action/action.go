package action

import (
	"crypto/rand"

	"github.com/arimotearipo/ggmp/internal/database"
	"github.com/arimotearipo/ggmp/internal/encryption"
	"github.com/arimotearipo/ggmp/internal/types"

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

func (a *Action) GetPassword(uri string) (username string, password string, err error) {
	username, encryptedPassword, err := a.db.GetPassword(uri, a.sess.id)
	if err != nil {
		return "", "", err
	}

	password, err = encryption.Decrypt(encryptedPassword, a.sess.masterKey)
	if err != nil {
		return "", "", err
	}

	return username, password, nil
}

func (a *Action) ListURIs() ([]types.URI, error) {
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
	encryptedPassword, err := encryption.Encrypt(password, a.sess.masterKey)
	if err != nil {
		return err
	}

	err = a.db.UpdatePassword(uri, username, encryptedPassword)
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

func (a *Action) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}

	err = a.db.AddMasterAccount(username, string(hashedPassword), salt)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) DeleteMasterAccount(username, password string) error {
	if err := a.Login(username, password); err != nil {
		return err
	}

	err := a.db.DeleteMasterAccount(username)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) ListMasterAccounts() ([]string, error) {
	accounts, err := a.db.ListMasterAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *Action) UpdateMasterPassword(newMasterPassword string) error {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}

	newMasterKey := encryption.DeriveKey([]byte(newMasterPassword), salt)

	savedOldMasterKey := make([]byte, len(a.sess.masterKey))
	copy(savedOldMasterKey, a.sess.masterKey)

	// get all the uris associated to the logged in master account
	uris, err := a.db.ListURIs(a.sess.id)
	if err != nil {
		return err
	}

	for _, uri := range uris {
		// get the decrypted password and username
		username, decryptedPassword, err := a.GetPassword(uri.Uri)
		if err != nil {
			return err
		}

		// temporarily assign to the new master key
		a.sess.masterKey = newMasterKey

		// call update. UpdatePassword() will re-encrypt the password with new master key
		err = a.UpdatePassword(uri.Uri, username, decryptedPassword)
		if err != nil {
			// if something goes wrong, use back old master key
			a.sess.masterKey = savedOldMasterKey
			return err
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newMasterPassword), 12)
	if err != nil {
		// TODO: handle reversing entire action
		return err
	}

	err = a.db.ChangeMasterPassword(a.sess.id, string(hashedPassword), salt)
	if err != nil {
		return err
	}

	return nil
}

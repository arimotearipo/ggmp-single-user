package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(name string) *Database {
	// connecting to database
	DB, err := sql.Open("sqlite", name)
	if err != nil {
		log.Fatal("Fail to connect to database")
	}
	log.Println("Connected to database")

	// creating master account table
	createMasterAccountSchema := `CREATE TABLE IF NOT EXISTS master_account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT UNIQUE,
		"hashed_password" TEXT,
		"initialization_vector" BLOB NOT NULL,
		"salt" BLOB NOT NULL
	);`
	statement, err := DB.Prepare(createMasterAccountSchema)
	if err != nil {
		log.Println("err:", err)
		log.Fatal("Fail to prepare createMasterAccountSchema SQL statement")
	}
	statement.Exec()
	log.Println("Master account table created")

	// create accounts table
	createAccountsSchema := `CREATE TABLE IF NOT EXISTS accounts (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"owner" REFERENCES master_account (id),
		"uri" TEXT,
		"username" TEXT,
		"encrypted_password" TEXT
	);`
	statement, err = DB.Prepare(createAccountsSchema)
	if err != nil {
		log.Println("err:", err)
		log.Fatal("Fail to prepare createAccountsSchema SQL statement")
	}
	statement.Exec()
	log.Println("Account table created")

	return &Database{DB}
}

func (db *Database) Close() {
	db.DB.Close()
}

func (db *Database) AddPassword(id int, uri, username, encryptedPassword string) error {
	insertAccountQuery := `INSERT INTO accounts (owner, uri, username, encrypted_password) VALUES (?, ?, ?, ?);`
	statement, err := db.DB.Prepare(insertAccountQuery)
	if err != nil {
		return err
	}

	statement.Exec(id, uri, username, encryptedPassword)
	return nil
}

func (db *Database) GetPassword(uri string, id int) (string, string, error) {
	selectAccountQuery := `SELECT username, encrypted_password FROM accounts WHERE uri = ? AND owner = ?;`
	statement, err := db.DB.Prepare(selectAccountQuery)

	if err != nil {
		return "", "", err
	}

	var username string
	var encryptedPassword string
	err = statement.QueryRow(uri, id).Scan(&username, &encryptedPassword)
	if err != nil {
		return "", "", err
	}

	return username, encryptedPassword, nil
}

func (db *Database) ListURIs(id int) ([]string, error) {
	selectAccountQuery := `SELECT uri FROM accounts WHERE owner = ?;`
	statement, err := db.DB.Prepare(selectAccountQuery)

	if err != nil {
		return nil, err
	}

	rows, err := statement.Query(id)

	if err != nil {
		return nil, err
	}

	var uris []string
	for rows.Next() {
		var uri string
		rows.Scan(&uri)
		uris = append(uris, uri)
	}
	return uris, nil
}

func (db *Database) DeleteAccount(uri string, id int) error {
	deleteAccountQuery := `DELETE FROM accounts WHERE uri = ? AND owner = ?;`
	statement, err := db.DB.Prepare(deleteAccountQuery)

	if err != nil {
		return err
	}

	_, err = statement.Exec(uri, id)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) UpdatePassword(uri string, username string, encryptedPassword string) error {
	updateAccountQuery := `UPDATE accounts SET username = ?, encrypted_password = ? WHERE uri = ?;`
	statement, err := db.DB.Prepare(updateAccountQuery)

	if err != nil {
		return err
	}

	statement.Exec(username, encryptedPassword, uri)
	return nil
}

func (db *Database) AddMasterAccount(username string, hashedPassword string, initializationVector, salt []byte) error {
	insertMasterAccountQuery := `INSERT INTO master_account (username, hashed_password, initialization_vector, salt) VALUES (?, ?, ?, ?);`
	statement, err := db.DB.Prepare(insertMasterAccountQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = statement.Exec(username, hashedPassword, initializationVector, salt)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetMasterAccount(username string) (string, []byte, []byte, error) {
	selectMasterAccountQuery := `SELECT COUNT(*) FROM master_account WHERE username = ?;`
	statement, err := db.DB.Prepare(selectMasterAccountQuery)
	if err != nil {
		fmt.Println(err)
		return "", nil, nil, err
	}

	var count int
	statement.QueryRow(username).Scan(&count)
	if count == 0 {
		return "", nil, nil, errors.New("username does not exist")
	}

	selectMasterAccountQuery = `SELECT hashed_password, initialization_vector, salt FROM master_account WHERE username = ?;`
	statement, err = db.DB.Prepare(selectMasterAccountQuery)
	if err != nil {
		fmt.Println(err)
		return "", nil, nil, err
	}

	var hashedPassword string
	var initializationVector, salt []byte
	statement.QueryRow(username).Scan(&hashedPassword, &initializationVector, &salt)
	return hashedPassword, initializationVector, salt, nil
}

func (db *Database) DeleteMasterAccount(username string) error {
	deleteLoginsQuery := `DELETE FROM accounts WHERE owner = (SELECT id FROM master_account WHERE username = ?);`
	statement, err := db.DB.Prepare(deleteLoginsQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(username)
	if err != nil {
		return err
	}

	deleteMasterAccountQuery := `DELETE FROM master_account WHERE username = ?;`
	statement, err = db.DB.Prepare(deleteMasterAccountQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(username)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ListMasterAccounts() ([]string, error) {
	selectMasterAccountQuery := `SELECT username FROM master_account;`
	statement, err := db.DB.Prepare(selectMasterAccountQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var usernames []string
	for rows.Next() {
		var username string
		rows.Scan(&username)
		usernames = append(usernames, username)
	}
	return usernames, nil
}

func (db *Database) GetUserId(username string) (int, error) {
	query := "SELECT id FROM master_account WHERE username = ?;"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	var id int
	err = statement.QueryRow(username).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

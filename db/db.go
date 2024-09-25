package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


type Database struct {
	DB *sql.DB
}

func NewDatabase(name string) *Database {
	DB, err := sql.Open("sqlite3", name)

	if err != nil {
		log.Fatal("Fail to connect to database")
	}

	log.Println("Connected to database")
	return &Database{DB}
}

func (d *Database) CreateSchema() {
	createAccountsSchema := `CREATE TABLE IF NOT EXISTS account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"hashed_password" TEXT
	);`

	statement, err := d.DB.Prepare(createAccountsSchema)

	if err != nil {
		log.Println("err:", err)
		log.Fatal("Fail to prepare createAccountsSchema SQL statement")
	}


	statement.Exec()
	log.Println("Account table created")
}

func (d *Database) RegisterAccount(u string, h string) {
	register := `INSERT INTO account VALUES ($1, $2);`

	_, err := d.DB.Exec(register, u, h)

	if err != nil {
		log.Println(err)
		log.Fatal("Fail to register account")
	}

	log.Println("Account registered")
}
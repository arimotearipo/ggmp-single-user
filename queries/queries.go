package queries

import (
	"database/sql"
	"log"

	"github.com/arimotearipo/ggmp/db"
)

type Queries struct {
	DB *sql.DB
}

func NewQueries(db *db.Database) *Queries {
	return &Queries{db.DB}
}

func (q *Queries) RegisterAccount(username string, hashedPassword string) {
	sql := `INSERT INTO account (username, hashed_password) VALUES ($1, $2);`

	_, err := q.DB.Exec(sql, username, hashedPassword)

	if err != nil {
		log.Println(err)
		log.Fatal("Fail to register account")
	}

	log.Println("Account registered")
}

func (q *Queries) SignInAccount(username string) string {
	sql := `SELECT (hashed_password) FROM account WHERE username = $1;`

	var hashedPassword string
	row := q.DB.QueryRow(sql, username)
	err := row.Scan(&hashedPassword)

	if err != nil {
		log.Println(err)
		log.Fatal("Fail to retrieve hashed password")
	}

	return hashedPassword
}
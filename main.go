package main

import (
	"github.com/arimotearipo/ggmp/cmd"
	"github.com/arimotearipo/ggmp/db"
	"github.com/arimotearipo/ggmp/queries"
)


func main() {
	database := db.NewDatabase("./ggmp.db")

	defer database.DB.Close()

	database.CreateSchema()

	queries := queries.NewQueries(database)

	cmd.SetQueries(queries)

	cmd.Execute()
}	
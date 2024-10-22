package main

import (
	"flag"
	"fmt"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/arimotearipo/ggmp/internal/database"
	teamodels "github.com/arimotearipo/ggmp/internal/tea_models"
	"github.com/arimotearipo/ggmp/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const ggmp string = "\n\n" + `      ::::::::   ::::::::    :::   :::   :::::::::
    :+:    :+: :+:    :+:  :+:+: :+:+:  :+:    :+:
   +:+        +:+        +:+ +:+:+ +:+ +:+    +:+ 
  :#:        :#:        +#+  +:+  +#+ +#++:++#+   
 +#+   +#+# +#+   +#+# +#+       +#+ +#+          
#+#    #+# #+#    #+# #+#       #+# #+#           
########   ########  ###       ### ###            ` + "\n\n"

const gogetmypassword string = "\t\tgo-get-my-password\n\n"

func main() {
	databaseFile := flag.String("file", "ggmp.db", "file to store passwords")

	flag.Parse()

	file := utils.FixExtension(*databaseFile)

	db := database.NewDatabase(file)
	defer db.Close()

	a := action.NewAction(db)

	defer a.EncryptDBFile() // encrypt db file before exiting

	fmt.Print(ggmp)
	fmt.Print(gogetmypassword)

	model := teamodels.NewAuthMenuModel(a)
	programme := tea.NewProgram(model)

	if _, err := programme.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

	defer utils.ClipboardWriteAndErase("", 0) // clear clipboard before programme exits
}

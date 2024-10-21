package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/arimotearipo/ggmp/internal/database"
	"github.com/arimotearipo/ggmp/internal/encryption"
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

func verifySQLite(databaseFile string) error {
	// check if databaseFile passed is actually a new file location
	_, err := os.Stat(databaseFile)
	if err != nil {
		return nil
	}

	file, err := os.Open(databaseFile)
	if err != nil {
		return err
	}
	defer file.Close()

	header := make([]byte, 16)
	_, err = file.Read(header)
	if err != nil {
		return err
	}

	if string(header) != "SQLite format 3\x00" {
		return errors.New("invalid SQLite file")
	}

	return nil
}

func decryptFile(databaseFile string, secretKey string) {
	if secretKey == "" {
		panic("Not a valid SQLite file\nMaybe it is encrypted? If yes, use the -key flag and pass the secret key")
	}

	err := encryption.DecryptFile(databaseFile, []byte(secretKey))
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseFile := flag.String("file", "ggmp.db", "file to store passwords")
	secretKey := flag.String("key", "", "secret to decrypt database file")

	flag.Parse()

	file := utils.FixExtension(*databaseFile)

	err := verifySQLite(file)
	if err != nil {
		decryptFile(*databaseFile, *secretKey)
	}

	db := database.NewDatabase(file)
	defer db.Close()

	a := action.NewAction(db)

	fmt.Print(ggmp)
	fmt.Print(gogetmypassword)

	model := teamodels.NewAuthMenuModel(a)
	programme := tea.NewProgram(model)

	if _, err := programme.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

	defer utils.ClipboardWriteAndErase("", 0) // clear clipboard before programme exits
}

package main

import (
	"fmt"

	"github.com/arimotearipo/ggmp/action"
	"github.com/arimotearipo/ggmp/database"
	teamodels "github.com/arimotearipo/ggmp/tea_models"
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

	db := database.NewDatabase("ggmp.db")
	defer db.Close()

	a := action.NewAction(db)

	fmt.Print(ggmp)
	fmt.Print(gogetmypassword)

	model := teamodels.NewAuthMenuModel(a)
	programme := tea.NewProgram(model)

	if _, err := programme.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}

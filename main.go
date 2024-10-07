package main

import (
	"fmt"

	"github.com/arimotearipo/ggmp/cmd"
	"github.com/arimotearipo/ggmp/database"
	teamodels "github.com/arimotearipo/ggmp/tea_models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	fmt.Println("Welcome to GGMP CLI")

	db := database.NewDatabase("ggmp.db")
	defer db.Close()

	commands := cmd.NewCommands(db)

	model := teamodels.NewAuthMenuModel(commands)
	programme := tea.NewProgram(model)

	if _, err := programme.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}

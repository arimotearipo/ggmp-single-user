package main

import "fmt"

func readCommands() {
	for {
		fmt.Println("1. Retrieve password")
		fmt.Println("2. Add password")
		fmt.Println("3. Exit")
		fmt.Printf("Enter your choice: ")

		var choice int
		fmt.Scanf("%d", &choice)
		fmt.Println(choice)

		switch choice {
		case 1:
			fmt.Println("Retrieve password")
		case 2:
			fmt.Println("Add password")
		case 3:
			fmt.Println("Exit")
			return
		}
	}
}

func main() {
	fmt.Println("Welcome to GGMP CLI")

	readCommands()
}

package main

import (
	"bufio"
	"fmt"
	"menu-app/menu"
	"os"
	"strings"
)

var in *bufio.Reader

func main() {
	fmt.Println("Welcome to the Menu App!")
	// create buffered reader to read user input
	in = bufio.NewReader(os.Stdin)
	choice := "a" // initial choice invalid

	for choice != "q" {
		presentOptions()
		choice, _ = in.ReadString('\n')
		choice = strings.TrimSpace(choice)
		processOptions(choice)
	}
}

func processOptions(choice string) {
	switch choice {
	case "1":
		menu.Print()
	case "2":
		err := menu.AddItem(in)
		if err != nil {
			fmt.Println(fmt.Errorf("unable to add item: %w", err))
		}
	case "q":
		fmt.Println("Exiting Application")
	default:
		fmt.Println("Please enter a valid option")
	}
}

func presentOptions() {
	fmt.Println("Please Select an option to continue:")
	fmt.Println("1) Print Menu")
	fmt.Println("2) Add Item")
	fmt.Println("q) Quit Application")
}

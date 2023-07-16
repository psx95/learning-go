package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Menu App!")
	// create buffered reader to read user input
	in := bufio.NewReader(os.Stdin)
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
		fmt.Println("Will Print Menu")
	case "2":
		fmt.Println("Will Add Items")
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

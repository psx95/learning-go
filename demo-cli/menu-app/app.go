package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MenuItem struct {
	name   string
	prices map[string]float64
}

// Initialize menu with some pre-filled entries
var menu = []MenuItem{
	{name: "Coffee", prices: map[string]float64{"venti": 5.45, "grande": 3.50, "tall": 2.20}},
	{name: "Hot Chocolate", prices: map[string]float64{"large": 3.00, "medium": 2.50, "small": 2.00}},
}

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
		printMenu()
	case "2":
		addMenuItem()
	case "q":
		fmt.Println("Exiting Application")
	default:
		fmt.Println("Please enter a valid option")
	}
}

func printMenu() {
	for _, menuItem := range menu {
		fmt.Println(menuItem.name)
		fmt.Println(strings.Repeat("-", 20))
		for size, price := range menuItem.prices {
			fmt.Printf("%20s\t%.2f\n", size, price)
		}
	}
}

func addMenuItem() {
	fmt.Println("Enter name of Menu Item")
	newItemName, err := in.ReadString('\n')
	if err == nil {
		newItem := MenuItem{
			name:   strings.TrimSpace(newItemName),
			prices: map[string]float64{},
		}
		menu = append(menu, newItem)
	} else {
		fmt.Println("Error accpeting menu item name", err)
	}
}

func presentOptions() {
	fmt.Println("Please Select an option to continue:")
	fmt.Println("1) Print Menu")
	fmt.Println("2) Add Item")
	fmt.Println("q) Quit Application")
}

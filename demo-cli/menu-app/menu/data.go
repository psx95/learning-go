package menu

import (
	"bufio"
	"fmt"
	"strings"
)

// Initialize menu with some pre-filled entries
var menuItems = []MenuItem{
	{name: "Coffee", prices: map[string]float64{"venti": 5.45, "grande": 3.50, "tall": 2.20}},
	{name: "Hot Chocolate", prices: map[string]float64{"large": 3.00, "medium": 2.50, "small": 2.00}},
}

func AddMenuItem(inputStream *bufio.Reader) {
	fmt.Println("Enter name of Menu Item")
	newItemName, err := inputStream.ReadString('\n')
	if err == nil {
		newItem := MenuItem{
			name:   strings.TrimSpace(newItemName),
			prices: map[string]float64{},
		}
		menuItems = append(menuItems, newItem)
	} else {
		fmt.Println("Error accpeting menu item name", err)
	}
}

func GetCurrentMenuItems() []MenuItem {
	return menuItems
}

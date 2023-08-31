package menu

import (
	"bufio"
	"fmt"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, menuItem := range m {
		fmt.Println(menuItem.name)
		fmt.Println(strings.Repeat("-", 20))
		for size, price := range menuItem.prices {
			fmt.Printf("%20s\t%.2f\n", size, price)
		}
	}
}

func (m *menu) addItem(inputStream *bufio.Reader) {
	fmt.Println("Enter name of Menu Item")
	newItemName, err := inputStream.ReadString('\n')
	if err == nil {
		newItem := menuItem{
			name:   strings.TrimSpace(newItemName),
			prices: map[string]float64{},
		}
		*m = append(*m, newItem)
	} else {
		fmt.Println("Error accpeting menu item name", err)	
	}
}

func Print() {
	data.print()
}

func AddItem(inputStream *bufio.Reader) {
	data.addItem(inputStream)
}

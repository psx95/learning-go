package menu

import (
	"bufio"
	"errors"
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

func (m *menu) addItem(inputStream *bufio.Reader) error {
	fmt.Println("Enter name of Menu Item")
	newItemName, err := inputStream.ReadString('\n')
	newItemName = strings.TrimSpace(newItemName)
	isDuplicate := m.checkDuplicate(newItemName)
	if isDuplicate {
		return errors.New("invalid input: Duplicate item")
	} else {
		fmt.Println("Not duplicate")
	}
	if err == nil {
		newItem := menuItem{
			name:   strings.TrimSpace(newItemName),
			prices: map[string]float64{},
		}
		*m = append(*m, newItem)
	} else {
		fmt.Println("Error accpeting menu item name", err)
		return err
	}
	return nil
}

func (m *menu) checkDuplicate(newItem string) bool {
	for _, item := range *m {
		if item.name == newItem {
			return true
		}
	}
	return false
}

func Print() {
	data.print()
}

func AddItem(inputStream *bufio.Reader) error {
	return data.addItem(inputStream)
}

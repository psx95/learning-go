package menu

import (
	"fmt"
	"strings"
)

type MenuItem struct {
	name   string
	prices map[string]float64
}

func PrintCurrentMenu() {
	for _, menuItem := range GetCurrentMenuItems() {
		fmt.Println(menuItem.name)
		fmt.Println(strings.Repeat("-", 20))
		for size, price := range menuItem.prices {
			fmt.Printf("%20s\t%.2f\n", size, price)
		}
	}
}

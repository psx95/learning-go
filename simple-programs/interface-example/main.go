package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	name string
	id   int64
}

// Implementing printer interface on the user type
func (u user) Print() string {
	return fmt.Sprintf("%v [%v]", u.name, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

// Implementing printer interface on the menuItem type
func (m menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(m.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range m.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func main() {
	// Declare interface
	var p printer
	p = user{
		name: "psx",
		id:   1995,
	}
	fmt.Println(p.Print())

	p = menuItem{
		name: "Coffee",
		prices: map[string]float64{
			"small":  1.65,
			"medium": 2.25,
			"large":  2.75,
		},
	}
	fmt.Println(p.Print())

	fmt.Println("Checking concrete types using Type Assertions")
	// Type Assertions - Checking if p is of type user
	u, ok := p.(user)
	fmt.Println(u, ok)

	// Type Assertiosn - Checking if p is of type menuItem
	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	fmt.Println("Checking concrete types using Type Switch")

	switch v := p.(type) {
	case user:
		fmt.Printf("The underlying type is user: [%s: %d]\n", v.name, v.id)
	case menuItem:
		fmt.Println("The underlying type is menuItem:", v)
	default:
		fmt.Printf("Cannot determine underlying type of printer: %s", v.Print())
	}

}

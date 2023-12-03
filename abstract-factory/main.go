// Resource: https://refactoring.guru/design-patterns/abstract-factory/go/example#example-0--adidasShoe-go

package main

import (
	"fmt"
	"os"
)

func displayShirt(s IShirt) {
	fmt.Println("Your Shirt")
	fmt.Printf("Logo: %s Length: %d\n", s.getLogo(), s.getLength())
}

func displayShoe(s IShoe) {
	fmt.Println("Your Shoe")
	fmt.Printf("Logo: %s Size: %d\n", s.getLogo(), s.getSize())
}

func displayChoices() string {
	var userInput string
	fmt.Println("Welcome to ABC Sports Authority")
	fmt.Println("Type brand name to buy your set")
	fmt.Println("1. adidas")
	fmt.Println("2. nike")
	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return userInput
}

func main() {
	for {
		choice := displayChoices()
		sportsFactory, err := GetSportsFactory(choice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		shoe := sportsFactory.makeShoe()
		shirt := sportsFactory.makeShirt()
		displayShoe(shoe)
		displayShirt(shirt)

	}
}

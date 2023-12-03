package main

import "fmt"

func main() {
	ak47, _ := createGun("ak47")
	rifle, _ := createGun("rifle")

	fmt.Printf("Gun: %s | Power: %.2f\n", ak47.GetName(), ak47.Power())
	fmt.Printf("Gun: %s | Power: %.2f\n", rifle.GetName(), rifle.Power())
}

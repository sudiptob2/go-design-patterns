package main

import "fmt"

func createGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "rifle" {
		return newRifle(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

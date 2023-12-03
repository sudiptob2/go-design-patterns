package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		// return adidas concrete factory
		return &AdidasFactory{}, nil
	}
	if brand == "nike" {
		return &NikeFactory{}, nil
	}
	return nil, fmt.Errorf("brand is not supported")
}

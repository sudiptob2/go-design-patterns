package main

type IShirt interface {
	setLogo(logo string)
	setLength(length int)
	getLogo() string
	getLength() int
}

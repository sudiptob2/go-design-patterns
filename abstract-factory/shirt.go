package main

type Shirt struct {
	length int
	logo   string
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setLength(length int) {
	s.length = length
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getLength() int {
	return s.length
}

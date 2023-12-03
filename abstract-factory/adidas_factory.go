package main

type AdidasFactory struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *AdidasFactory) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 7,
		},
	}
}

func (a *AdidasFactory) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo:   "adidas",
			length: 14,
		},
	}
}

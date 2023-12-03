package main

type NikeFactory struct {
}

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (nf *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 3,
		},
	}
}

func (nf *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo:   "nike",
			length: 26,
		},
	}
}

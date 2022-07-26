package main

type Nike struct {
}

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 0,
		}}
}

func (a *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 0,
		}}
}

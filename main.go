package main

import (
	"fmt"
	"log"
)

func main() {
	//Imran:= Nike{NikeShirt{Shirt{
	//	logo: "",
	//	size: 0,
	//}}}
	//Imran.makeShirt()
	adidasFactory, err := GetSportsFactory("adidas")
	if err != nil {
		log.Println("Error Occured", err.Error())
	}

	nikeFactory, err := GetSportsFactory("nike")
	if err != nil {
		log.Println("Error Occured", err.Error())
	}

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	PrintShoeDetails(adidasShoe)
	PrintShirtDetails(adidasShirt)

	PrintShoeDetails(nikeShoe)
	PrintShirtDetails(nikeShirt)
}

func PrintShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s ", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func PrintShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s ", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

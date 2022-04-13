package main

import "fmt"

// automatically called
func main() {
	cards := []string{newCard(), newCard()}

	//Does not modify the slice, assign new slice.
	cards = append(cards, "Six of Spades")

	// range return idx and element
	// redeclare variable whenever it's iterated.
	// declare but not used make compile failed
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	// return type is placed behind.
	return "Five of Diamonds"
}

// As long as, files are in same package, it's able to use the thing in the other.
// Array vs Slice

// Slice is an array ( fancy name ) , more func, can grow & shrink
// Every element in a slice must be of same type

// Array primitve data type fixed length

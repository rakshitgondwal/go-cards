package main

import "fmt"

type deck [] string 

func newDeck() deck{
	cards := []string{}

	cardSuits:= deck{"Spades", "Hearts", "Diamonds"}
	cardValues:= deck{"Ace","Two","Three"}

	for _,suit:= range cardSuits{
		for	_,value:= range cardValues{
			cards = append(cards, value +" of "+suit)
		}
	}

	return cards
}

func (d deck) print(){
	for _,card := range d{
		fmt.Println(card)
	}
}
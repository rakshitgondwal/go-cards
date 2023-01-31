package main

import (
	"fmt"
	"math/rand"
	"time"
)
type deck [] string 

func newDeck() deck{
	cards := deck{}

	cardSuits:= []string{"Spades", "Hearts", "Diamonds"}
	cardValues:= []string{"Ace","Two","Three"}

	for _,suit:= range cardSuits{
		for	_,value:= range cardValues{
			cards = append(cards, value +" of "+suit)
		}
	}

	return cards
}


func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

func (d deck) print(){
	for _,card := range d{
		fmt.Println(card)
	}
}

func (d deck) shuffle(){

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newpostion := r.Intn(len(d) -1)

		d[i], d[newpostion] = d[newpostion], d[i]
	}
}

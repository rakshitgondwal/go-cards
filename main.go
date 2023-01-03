package main


func main(){
	cards := newDeck()
	leftCards, _ := deal(cards,2)
	leftCards.print();
}

func newCard() string{
	return "Ace Of Diamonds"
}
package main


func main(){
	cards := deck{newCard(), "Five of Spades", "Three of Hearts"}
	cards = append(cards, "King of Diamond")

	cards.print();
}

func newCard() string{
	return "Ace Of Diamonds"
}
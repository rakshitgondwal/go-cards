package main
func main(){
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
func newCard() string{
	return "Ace Of Diamonds"
};

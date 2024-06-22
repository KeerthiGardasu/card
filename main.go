package main

func main() {
	cards := newdeck()
	//hand, remainingcards := deal(cards, 5)
	//hand.print()
	//remainingcards.print()
	//fmt.Println(cards.tostring())
	cards.saveToFile("my_cards")

}

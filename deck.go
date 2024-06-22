package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newdeck() deck {
	cards := deck{}
	Cardsuits := []string{"hearts", "spades", "club", "diamonds"}
	Cardvalues := []string{"Ace", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, suit := range Cardsuits {
		for _, value := range Cardvalues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]

}
func (d deck) tostring() string {
	return strings.Join([]string(d), ",")

}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.tostring()), 0666)
}

package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var cards = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var isBefore = map[string] bool{ 
	"2":false,
	"3":false,
	"4":false,
	"5":false,
	"6":false,
	"7":false,
	"8":false,
	"9":false,
	"10":false,
	"J":false,
	"Q":false,
	"K":false,
	"A":false,
}
var countCards int = 0
func getValue(card string) int {

	if card == "J" {
		return 11
	}
	if card == "Q" {
		return 12
	}
	if card == "K" {
		return 13
	}
	if card == "A" {
		return 14
	}
	value, _ := strconv.ParseInt(card, 10, 0)
	return int(value)

}

func Something(num int) int {
	return num * num
}

func Randomiser() string {
	if countCards == len(cards){
		return "Cards Exhausted"
	}
	rand.Seed(time.Now().Unix())
	var my_card string = cards[rand.Intn(len(cards))]
	for isBefore[my_card] {
		my_card = cards[rand.Intn(len(cards))]
	}
	isBefore[my_card] = true
	countCards++
	return my_card
}

func IsAbove(prevCard string, currentCard string) bool {
	return getValue((currentCard)) > getValue(prevCard)
}
func IsEqual(prevCard string, currentCard string) bool {
	return getValue((currentCard)) == getValue(prevCard)
}

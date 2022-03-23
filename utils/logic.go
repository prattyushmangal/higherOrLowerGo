package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var cards = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

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
	rand.Seed(time.Now().Unix())

	return cards[rand.Intn(len(cards))]
}

func IsAbove(prevCard string, currentCard string) bool {
	return getValue((currentCard)) > getValue(prevCard)
}
func IsEqual(prevCard string, currentCard string) bool {
	return getValue((currentCard)) == getValue(prevCard)
}

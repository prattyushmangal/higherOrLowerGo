package utils

import (
	"math/rand"
	"time"
)

var cards = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func Something(num int) int {
	return num * num
}

func Randomiser() string {
	rand.Seed(time.Now().Unix())
	return cards[rand.Intn(len(cards))]
}

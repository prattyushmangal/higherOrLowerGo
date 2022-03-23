package utils_test

import (
	logic "anush22/cardgame/utils"
	"testing"

	"github.com/stretchr/testify/assert"
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


func TestSomething(t *testing.T) {
	funcOutput := logic.Something(9)
	assert.Equal(t, funcOutput, 81)
}

func TestRandomiser(t *testing.T) {
	
	for i := 1; i <= 10; i++{
	funcOutput := logic.Randomiser()
	assert.Contains(t, cards, funcOutput) 
	assert.False(t, isBefore[funcOutput])
	isBefore[funcOutput]= true
	}
}

func TestIsAboveSimple(t *testing.T) {
	previousCard := "2"
	currentCard := "A"
	funcOutput := logic.IsAbove(previousCard, currentCard)
	assert.True(t, funcOutput)
}

func TestIsAboveNotTrue(t *testing.T) {
	previousCard := "A"
	currentCard := "2"
	funcOutput := logic.IsAbove(previousCard, currentCard)
	assert.False(t, funcOutput)
}
func TestIsEqual(t *testing.T) {
	previousCard := "A"
	currentCard := "2"
	funcOutput := logic.IsEqual(previousCard, currentCard)
	assert.False(t, funcOutput)
}

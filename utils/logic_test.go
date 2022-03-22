package utils_test

import (
	logic "anush22/cardgame/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cards = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func TestSomething(t *testing.T) {
	funcOutput := logic.Something(9)
	assert.Equal(t, funcOutput, 81)
}

func TestRandomiser(t *testing.T) {
	funcOutput := logic.Randomiser()
	assert.Contains(t, cards, funcOutput)
}

package src_test

import (
	src "anush22/cardgame/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = [4]string{"A", "B", "a", "b"}

func TestGetInput(t *testing.T) {
	funcOutput := src.GetInput()
	assert.Contains(t, options, funcOutput)
}

// Test Running Game
// - Test Outputting Card
// - Test Taking Input
// - Test assertion based on input
// - Test Ending Game on Failed Input
// Test Point Keeping
// Point Showing
// Test Restarting game
// https://github.com/faiface/pixel & https://www.codementor.io/@ajinkyax/game-development-in-golang-15yxglh1kxwsl

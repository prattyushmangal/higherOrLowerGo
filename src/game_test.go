package src_test

import (
	game "anush22/cardgame/src"
	"testing"
	logic "anush22/cardgame/utils"
	"github.com/stretchr/testify/assert"
)

var options = [2]string{"A", "B"}

// func TestGetInput(t *testing.T) {
// 	funcOutput := src.GetInput()
// 	assert.Contains(t, options, funcOutput)
// }

func TestUpdateScore(t *testing.T) {
	userAnswer:="A"
	prevScore:= game.GetScore()
	prevCard:="4"
	currentCard:="6"
	game.UpdateScore(prevCard,  currentCard, userAnswer)
	funcOutput:= game.GetScore()
	assert.GreaterOrEqual(t, funcOutput, 0)
	assert.LessOrEqual(t, funcOutput, 12)
	assert.GreaterOrEqual(t, prevScore, 0)
	assert.LessOrEqual(t,  prevScore, 12)
	if (userAnswer=="A" && logic.IsAbove(prevCard, currentCard)){
		assert.Greater(t,  funcOutput, prevScore)
	} else if (userAnswer=="B" && !logic.IsAbove(prevCard, currentCard)){
		assert.Greater(t,  funcOutput, prevScore)
	} else{
		assert.Equal(t, prevScore, funcOutput)
	}
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

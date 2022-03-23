package main

import (
	logic "anush22/cardgame/utils"
	game "anush22/cardgame/src"
	"fmt"
)

func main() {
	// fmt.Println("Hello, world.")
	// squared := logic.Something(9)
	// fmt.Println(squared)
	// fmt.Println(logic.Randomiser())
	// answer:=game.GetInput()
	// fmt.Println(answer)

	var guess string
	var prevCard string
	var currentCard string
	var game_score int
	game_score=0
	fmt.Println("Game Starts")
	fmt.Println("Your score is " ,game_score)
	fmt.Println("First card is : ")
	prevCard =logic.Randomiser()
	fmt.Println(prevCard)
	for true{
		currentCard = logic.Randomiser()
		if(currentCard == "Cards Exhausted"){
			fmt.Println("Congratulations. You Win. Your score is " , game_score)
			break
		}
		fmt.Println("Time for next card")
		fmt.Println("Enter your Guess : A or B")
		fmt.Scanln(&guess)		
		fmt.Println("Your next card is : ")		
		fmt.Println(currentCard)
		game.UpdateScore(prevCard, currentCard, guess)
		if(game.GetScore() == game_score){
			fmt.Println("Sorry, Wrong Guess. Game ends. Your score is " ,game_score)
			break;
		}
		game_score++
		fmt.Println("Wow! That is a Correct Guess. Your score is " ,game_score)
		prevCard=currentCard


	}
}

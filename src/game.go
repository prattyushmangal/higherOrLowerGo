package src

import (
	
	logic "anush22/cardgame/utils"
)

var options = [2]string{"A", "B"}
var prevCard string = ""
var currentCard string = ""
var score = 0
// func GetInput() string {
// 	var userInput string
// 	fmt.Scanln(&userInput)
// 	return userInput
// }

func UpdateScore(prevCard string,  currentCard string, userAnswer string){
	if (userAnswer=="A" && logic.IsAbove(prevCard, currentCard)){
		score++
	} else if (userAnswer=="B" && !logic.IsAbove(prevCard, currentCard)){
		score++
	}	
}
func GetScore() int{
	return score
}


package src

import "fmt"

var options = [4]string{"A", "B", "a", "b"}

func GetInput() string {
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

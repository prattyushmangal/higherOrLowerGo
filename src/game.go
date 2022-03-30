package src

import (
	
	logic "anush22/cardgame/utils"
	"fmt"
	"image"
	"os"
	// "math/rand"
	// "math"
	// "time"
	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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



func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func getSprite(card string) (*pixel.Sprite){
	path:="images/card"+card+".png"
	cardPicture, err := loadPicture(path)
	if err != nil {
		panic(err)
	}
	cardSprite := pixel.NewSprite(cardPicture, cardPicture.Bounds())
	return cardSprite
}

func Run(){
	cfg := pixelgl.WindowConfig{
		Title:  "Card Guess Game!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	
	var(
		prevCard string
		currentCard string
		game_score int
	)
	game_score=0
	prevmat := pixel.IM
	prevmat = prevmat.Moved(pixel.V(256, 384))
	currmat := pixel.IM
	currmat = currmat.Moved(pixel.V(768, 384))
	gameOverPicture, err := loadPicture("images/gameover.png")
	gameOverSprite := pixel.NewSprite(gameOverPicture, gameOverPicture.Bounds())

	prevCard = "Blank"
	currentCard =logic.Randomiser()
	prevSprite:=getSprite(prevCard)
	currSprite:=getSprite(currentCard)


	for !win.Closed() {


		if win.JustPressed(pixelgl.KeyDown) {
			prevCard = currentCard
			currentCard = logic.Randomiser()
			if(currentCard=="Cards Exhausted"){
				gameOverSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				break;
			}
			prevSprite = getSprite(prevCard)
			currSprite=getSprite(currentCard)
			UpdateScore(prevCard, currentCard, "B")
			if(GetScore() == game_score){
				fmt.Println("Sorry, Wrong Guess. Game ends. Your score is " ,game_score)
				fmt.Println(prevCard)
				fmt.Println(currentCard)
				gameOverSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				break;
			}
			game_score++
			fmt.Println("Wow! That is a Correct Guess. Your score is " ,game_score)
			fmt.Println(prevCard)
			fmt.Println(currentCard)
			
		}
		if win.JustPressed(pixelgl.KeyUp) {
			prevCard = currentCard
			currentCard = logic.Randomiser()
			if(currentCard=="Cards Exhausted"){
				gameOverSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				break;
			}
			prevSprite = getSprite(prevCard)
			currSprite=getSprite(currentCard)
			UpdateScore(prevCard, currentCard, "A")
			if(GetScore() == game_score){
				fmt.Println("Sorry, Wrong Guess. Game ends. Your score is " ,game_score)
				fmt.Println(prevCard)
				fmt.Println(currentCard)
				gameOverSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
				break;
			}
			game_score++
			fmt.Println("Wow! That is a Correct Guess. Your score is " ,game_score)
			fmt.Println(prevCard)
			fmt.Println(currentCard)
			
		}

		win.Clear(colornames.Greenyellow)

		prevSprite.Draw(win, prevmat)
		currSprite.Draw(win, currmat)
		win.Update()


	}
}


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


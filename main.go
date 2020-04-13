package main

// Need to get github.com/faiface/glhf
// and github.com/go-gl/glfw/v3.2/glfw
//installed locally before running this code

import (
	"fmt"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	_ "image/png"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Direction : the direction our hero will face (int)
type Direction int

// Size for our game window
const (
	windowHeight = 800
	windowWith   = 800
)

// Directions our hero can move
const (
	up Direction = iota
	down
	left
	right
)

// defaultRate: Default speed at which animation will be updated for a sprite
const defaultRate float64 = 1 / 5.0

//loadBoard: Main function to load game objects (ghosts, world setup
// in the wolrd map and ready to be drawn
func loadBoard(pm *pacman, sheet pixel.Picture, brd *board) {
	gh1 := &ghost{gridX: 5, gridY: 10, rate: defaultRate, spriteRow: 0, spriteCol: 0}
	gh2 := &ghost{gridX: 15, gridY: 14, rate: defaultRate, spriteRow: 1, spriteCol: 0}
	gh3 := &ghost{gridX: 8, gridY: 3, rate: defaultRate, spriteRow: 3, spriteCol: 0}
	gh4 := &ghost{gridX: 2, gridY: 5, rate: defaultRate, spriteRow: 1, spriteCol: 8}

	objectsToLoad := []loadable{brd, pm, gh1, gh2, gh3, gh4}
	for _, object := range objectsToLoad {
		err := object.load(sheet)
		if err != nil {
			panic(err)
		}
	}
	World.pm = pm
	World.brd = brd
	World.worldMap = worldMapLvl
	World.ghost = []*ghost{gh1, gh2, gh3, gh4}
	World.worldCoinCount = -1
}

//gameOver: Detects if our hero was caught by a ghost or if it ate
// all the coins. If so prints game over & the final score and returns true
func gameOver(basicAtlas *text.Atlas, win *pixelgl.Window) bool {
	if World.gameOver == true || World.score == World.worldCoinCount {
		time.Sleep(500 * time.Millisecond)
		basicTxt := text.New(pixel.V(300, 500), basicAtlas)
		_, err := fmt.Fprintln(basicTxt, "Game Over!")
		if err != nil {
			panic(err)
		}
		_, err2 := fmt.Fprintln(basicTxt, "Score: "+strconv.Itoa(World.score))
		if err2 != nil {
			panic(err2)
		}
		win.Clear(colornames.Black)
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4))
		win.Update()
		time.Sleep(3000 * time.Millisecond)
		return true
	} else {
		return false
	}
}

//displayScore: Prints current score on top right of the screen
func displayScore(basicAtlas *text.Atlas, win *pixelgl.Window) {
	basicTxt := text.New(pixel.V(500, 750), basicAtlas)
	_, err := fmt.Fprintln(basicTxt, "Score: "+strconv.Itoa(World.score))
	if err != nil {
		panic(err)
	}
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 3))
}

//getUserRequestedDirection: Determine if the user has pressed an arrow key
// and update the direction it is facing accordingly
func getUserRequestedDirection(win *pixelgl.Window) (Direction, bool) {
	direction := right
	userChanged := false
	if win.Pressed(pixelgl.KeyLeft) {
		direction = left
		userChanged = true
	} else if win.Pressed(pixelgl.KeyRight) {
		direction = right
		userChanged = true
	} else if win.Pressed(pixelgl.KeyUp) {
		direction = up
		userChanged = true
	} else if win.Pressed(pixelgl.KeyDown) {
		direction = down
		userChanged = true
	}
	return direction, userChanged
}

//Main function in the game.
func run() {

	//default facing direction for sprites
	direction := right

	//Create a window to run the game
	cfg := pixelgl.WindowConfig{
		Title:  "goPac!",
		Bounds: pixel.R(0, 0, windowWith, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//Load sprite sheet
	sheet, err := getSheet("spritemap-384.png")

	// Create our hero
	pm := &pacman{gridX: 1, gridY: 1, rate: defaultRate}
	err = pm.load(sheet)
	if err != nil {
		panic(err)
	}

	// Setup drawing buffer
	imd := imdraw.New(sheet)

	//Create and load board
	brd := &board{}
	loadBoard(pm, sheet, brd)
	last := time.Now()

	// Setup font to write message in the game screen
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Main game loop
	for !win.Closed() {
		// Run game

		//See if our hero was caught by ghosts of ate all coins
		if gameOver(basicAtlas, win) {
			break
		}

		//Reset screen and get ready to draw the board
		time.Sleep(105 * time.Millisecond)
		dt := time.Since(last).Seconds()

		win.Clear(colornames.Black)
		imd.Clear()

		//Determine if user wants our hero to change direction
		newDirection, requested := getUserRequestedDirection(win)
		if requested {
			direction = newDirection
		}

		// Draw board to buffer
		dawErr := brd.draw(imd)
		if dawErr != nil {
			panic(dawErr)
		}

		//Update our hero state and draw it to the buffer
		pm.update(dt, direction)
		pm.draw(imd)

		//Update the ghosts state and draw it to the buffer
		for _, gh := range World.ghost {
			gh.update(dt)
			gh.draw(imd)
		}

		//Draw buffer to window, update the score and update the window
		imd.Draw(win)
		displayScore(basicAtlas, win)
		win.Update()
	}
}

// Initialize and run  program
func main() {
	pixelgl.Run(run)
}

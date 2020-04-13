package main

import "github.com/faiface/pixel"

// board: Holds the pixel.picture representation of the game board
type board struct {
	sheet pixel.Picture
}

// load: Loads a given pixel.Picture into the board sheet
func (brd *board) load(sheet pixel.Picture) error {
	brd.sheet = sheet
	return nil
}

// draw: Take the coins and blocks information stored in the World struct and draw them in the
// and draw them in board based on World.worldMap level information
func (brd *board) draw(t pixel.Target) error {
	var coinCount int
	worldMap := World.worldMap
	blkFrame := getFrame(24, 24, 0, 5)
	coinFrame := getFrame(12, 12, 16, 19)
	for i := 0; i < len(worldMap); i++ {
		for j := 0; j < len(worldMap[0]); j++ {
			if worldMap[i][j] == 0 {
				block{frame: blkFrame, gridX: i, gridY: j, sheet: brd.sheet}.draw(t)
			} else if worldMap[i][j] == 1 {
				coin{frame: coinFrame, gridX: i, gridY: j, sheet: brd.sheet}.draw(t)
				coinCount += 1
			}
		}
	}
	if World.worldCoinCount == -1 {
		World.worldCoinCount = coinCount
	}
	return nil
}

//loadable: Interface to load a pixelPicture)
type loadable interface {
	load(pixel.Picture) error
}

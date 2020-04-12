// Contains logic for defining and updating the game board
package main

import "github.com/faiface/pixel"

type board struct {
	sheet pixel.Picture
}

func (brd *board) load(sheet pixel.Picture) error {
	brd.sheet = sheet
	return nil
}

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

type loadable interface {
	load(pixel.Picture) error
}

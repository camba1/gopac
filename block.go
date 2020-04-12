// Contains logic for defining the obstacles (blocks) in the game
package main

import (
	"github.com/faiface/pixel"
)

type block struct {
	frame pixel.Rect
	sheet pixel.Picture
	gridX int
	gridY int
}

func (blk block) draw(t pixel.Target) {
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	sprite.Set(blk.sheet, blk.frame)
	//pos := getRectInGrid(windowWith, windowHeight, worldMapZeorLen, worlMapLen, blk.gridY, blk.gridX)
	pos := getRectInGrid(windowWith, windowHeight, len(World.worldMap[0]), len(World.worldMap), blk.gridY, blk.gridX)
	sprite.Draw(t, pixel.IM.ScaledXY(pixel.ZV, pixel.V(
		pos.W()/sprite.Frame().W(),
		pos.H()/sprite.Frame().H(),
	)).Moved(pos.Center()),
	)
}

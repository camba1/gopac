// Contains logic for controlling placement of the coins
package main

import "github.com/faiface/pixel"

//coin: Structure to hold information about the coins drawn in the game
type coin struct {
	frame pixel.Rect
	sheet pixel.Picture
	gridX int
	gridY int
}

// draw: Determines the coin frame that needs to be loaded as well as where it needs to be drawn
func (cn coin) draw(t pixel.Target) {
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	sprite.Set(cn.sheet, cn.frame)
	pos := getRectInGrid(windowWith, windowHeight, len(World.worldMap[0]), len(World.worldMap), cn.gridY, cn.gridX)
	sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			pos.W()/sprite.Frame().W(),
			pos.H()/sprite.Frame().H(),
		)).
		Moved(pos.Center()),
	)
}

package main

import (
	"github.com/faiface/pixel"
	"math"
	"math/rand"
)

//ghost: Structure that holds all the information about the ghost to be drawn in the screen
type ghost struct {
	direction Direction                  //Direction the ghost is facing
	anims     map[Direction][]pixel.Rect //fames map
	rate      float64                    //animation rate
	frame     pixel.Rect                 //current frame
	sheet     pixel.Picture              //sprite sheet in pixel picture format
	gridX     int                        // X position in grid
	gridY     int                        // Y position in grid
	spriteRow int                        // Row for the Sprite
	spriteCol int                        // Column for the Sprite
}

// Load the animation frames to be used for each direction the ghost needs to move and loads it to the
// ghost.anims map
func (gh *ghost) load(sheet pixel.Picture) error {
	gh.sheet = sheet
	gh.direction = right
	gh.anims = make(map[Direction][]pixel.Rect)
	gh.frame = getFrame(24, 24, 1, 6)
	gh.anims[up] = append(gh.anims[up], getFrame(24, 24, gh.spriteCol+6, gh.spriteRow))
	gh.anims[up] = append(gh.anims[up], getFrame(24, 24, gh.spriteCol+7, gh.spriteRow))
	gh.anims[down] = append(gh.anims[down], getFrame(24, 24, gh.spriteCol+2, gh.spriteRow))
	gh.anims[down] = append(gh.anims[down], getFrame(24, 24, gh.spriteCol+3, gh.spriteRow))
	gh.anims[left] = append(gh.anims[left], getFrame(24, 24, gh.spriteCol+4, gh.spriteRow))
	gh.anims[left] = append(gh.anims[left], getFrame(24, 24, gh.spriteCol+5, gh.spriteRow))
	gh.anims[right] = append(gh.anims[right], getFrame(24, 24, gh.spriteCol+0, gh.spriteRow))
	gh.anims[right] = append(gh.anims[right], getFrame(24, 24, gh.spriteCol+1, gh.spriteRow))
	return nil
}

// draw: Determines the ghost frame that needs to be loaded as well as where it needs to be drawn
func (gh *ghost) draw(t pixel.Target) {
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	sprite.Set(gh.sheet, gh.frame)
	pos := getRectInGrid(windowWith, windowHeight, len(World.worldMap[0]), len(World.worldMap), gh.gridX, gh.gridY)
	sprite.Draw(t, pixel.IM.ScaledXY(
		pixel.ZV, pixel.V(
			pos.W()/sprite.Frame().W(),
			pos.H()/sprite.Frame().H(),
		)).Moved(pos.Center()),
	)
}

// Check if ghost if colliding with the game boundaries
func (gh *ghost) isCollidingWithWall(gridX, gridY int) bool {
	return gridX < 0 || gridX >= len(World.worldMap[0]) ||
		gridY < 0 || gridY > len(World.worldMap) ||
		World.worldMap[gridY][gridX] == 0

}

// Check if the ghost caught our hero
func (gh *ghost) isCollidingWithPacMan() bool {
	return World.pm.gridX == gh.gridX && World.pm.gridY == gh.gridY
}

// Determine the directions the ghost could move without running into a wall
// returns a slice with all the valid directions
func (gh *ghost) getPossibleMoves(oldGridx, oldGridy int) []Direction {
	gh.gridX = oldGridx
	gh.gridY = oldGridy
	possible := make([]Direction, 0)
	if World.worldMap[gh.gridY+1][gh.gridX] != 0 {
		possible = append(possible, up)
	}
	if World.worldMap[gh.gridY-1][gh.gridX] != 0 {
		possible = append(possible, down)
	}
	if World.worldMap[gh.gridY][gh.gridX+1] != 0 {
		possible = append(possible, right)
	}
	if World.worldMap[gh.gridY][gh.gridX-1] != 0 {
		possible = append(possible, left)
	}
	return possible
}

// update the driection the host is moving and detect if the game is over
func (gh *ghost) update(dt float64) {
	directionValue := gh.direction
	oldGridx := gh.gridX
	oldGridy := gh.gridY
	if directionValue == right {
		gh.gridX += 1
	} else if directionValue == left {
		gh.gridX -= 1
	} else if directionValue == up {
		gh.gridY += 1
	} else if directionValue == down {
		gh.gridY -= 1
	}
	if gh.isCollidingWithWall(gh.gridX, gh.gridY) {
		possibleMoves := gh.getPossibleMoves(oldGridx, oldGridy)
		gh.direction = possibleMoves[rand.Intn(len(possibleMoves))]
	}

	if gh.isCollidingWithPacMan() {
		World.gameOver = true
	}

	if gh.rate == 0 {
		gh.rate = 1
	}
	i := int(math.Floor(dt / gh.rate))
	if i == 0 {
		i = 1
	}
	gh.frame = gh.anims[gh.direction][i%len(gh.anims[gh.direction])]

}

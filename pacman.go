package main

import (
	"github.com/faiface/pixel"
	"math"
)

// Build our hero as a struct
type pacman struct {
	direction Direction                  //that our hero is facing
	anims     map[Direction][]pixel.Rect //list of frames for animation
	rate      float64                    // animation framerate
	frame     pixel.Rect                 // current frame
	sheet     pixel.Picture              // Sprite sheet reference
	gridX     int                        // current X position in grid
	gridY     int                        // current y position in grid
}

// Pacman functions

func (pm *pacman) load(sheet pixel.Picture) error {
	//var err error
	pm.sheet = sheet
	//if err != nil {
	//	panic(err)
	//}
	//Animation frame mapping. Note that 1,6 is the coords of image to be shown
	pm.anims = make(map[Direction][]pixel.Rect)
	pm.anims[up] = append(pm.anims[up], getFrame(24, 24, 1, 6))
	pm.anims[up] = append(pm.anims[up], getFrame(24, 24, 3, 6))
	pm.anims[down] = append(pm.anims[down], getFrame(24, 24, 5, 6))
	pm.anims[down] = append(pm.anims[down], getFrame(24, 24, 7, 6))
	pm.anims[left] = append(pm.anims[left], getFrame(24, 24, 0, 6))
	pm.anims[left] = append(pm.anims[left], getFrame(24, 24, 2, 6))
	pm.anims[right] = append(pm.anims[right], getFrame(24, 24, 4, 6))
	pm.anims[right] = append(pm.anims[right], getFrame(24, 24, 6, 6))

	return nil
}

func (pm *pacman) getNewGridPosition(direction Direction) (int, int) {
	if direction == right {
		return pm.gridX + 1, pm.gridY
	} else if direction == left {
		return pm.gridX - 1, pm.gridY
	} else if direction == up {
		return pm.gridX, pm.gridY + 1
	} else if direction == down {
		return pm.gridX, pm.gridY - 1
	} else {
		return pm.gridX, pm.gridY
	}
}

func (pm *pacman) draw(target pixel.Target) {
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	sprite.Set(pm.sheet, pm.frame)
	position := getRectInGrid(windowWith, windowHeight, 20, 20, pm.gridX, pm.gridY)
	sprite.Draw(target, pixel.IM.ScaledXY(pixel.ZV, pixel.V(
		position.W()/sprite.Frame().W(),
		position.H()/sprite.Frame().H(),
	)).Moved(position.Center()))
}

// update Pacman position and frame to be displayed
func (pm *pacman) update(dt float64, direction Direction) {
	newGridX, newGridY := pm.getNewGridPosition(direction)
	//oldGridX, oldGrdY := pm.gridX, pm.gridY
	if !pm.isCollidingWithWall(newGridX, newGridY) {
		pm.direction = direction
		pm.gridX, pm.gridY = newGridX, newGridY
	} else {
		newGridX, newGridY = pm.getNewGridPosition(pm.direction)
		if !pm.isCollidingWithWall(newGridX, newGridY) {
			pm.gridX, pm.gridY = newGridX, newGridY
		}
	}

	if pm.isCollidingWithGhost() {
		//pm.gridX, pm.gridY = oldGridX, oldGrdY
		World.gameOver = true
	}

	pm.eatingCoinCheck()

	if pm.rate == 0 {
		pm.rate = 1
	}
	i := int(math.Floor(dt / pm.rate))
	if i == 0 {
		i = 1
	}
	pm.frame = pm.anims[pm.direction][i%len(pm.anims[pm.direction])]

}

func (pm *pacman) isCollidingWithWall(gridX, gridY int) bool {
	return gridX < 0 || gridX >= len(World.worldMap[0]) ||
		gridY < 0 || gridY > len(World.worldMap) ||
		World.worldMap[gridY][gridX] == 0

}

func (pm *pacman) isCollidingWithGhost() bool {
	for _, gh := range World.ghost {
		if pm.gridX == gh.gridX && pm.gridY == gh.gridY {
			return true
		}
	}
	return false
}

func (pm *pacman) eatingCoinCheck() {
	if World.worldMap[pm.gridY][pm.gridX] == 1 {
		World.worldMap[pm.gridY][pm.gridX] = 2
		World.score++
	}
}

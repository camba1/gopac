package main

import "testing"

/*
	TestGetNewGridPosition: Ensure that our hero moves properly in the grid given a direction change
*/
func TestGetNewGridPosition(t *testing.T) {
	var newX, newY int
	pm := &pacman{gridX: 2, gridY: 2, rate: defaultRate}
	data := []struct {
		direction Direction
		gridX     int
		gridY     int
	}{
		{right, 3, 2},
		{left, 1, 2},
		{up, 2, 3},
		{down, 2, 1},
	}

	for _, move := range data {
		newX, newY = pm.getNewGridPosition(move.direction)
		if newX != move.gridX || newY != move.gridY {
			t.Errorf("Incorrect X, Y position for direction %d. Expected (%d,%d) and got (%d,%d)", move.direction, move.gridX, move.gridY, newX, newY)
		}

	}
}

/*
	Ensure our hero cannot go past the walls in the game
*/
func TestIsCollidingWithWall(t *testing.T) {
	World.worldMap = worldMapLvl
	pm := &pacman{gridX: 2, gridY: 2, rate: defaultRate}
	data := []struct {
		gridX       int
		gridY       int
		isCollision bool
	}{
		{gridX: 1, gridY: 1, isCollision: false},
		{len(worldMapLvl[0]) - 1, len(worldMapLvl) - 1, true},
		{len(worldMapLvl[0]) - 1, 2, true},
		{2, len(worldMapLvl) - 1, true},
	}

	for _, location := range data {
		collision := pm.isCollidingWithWall(location.gridX, location.gridY)
		if collision != location.isCollision {
			t.Errorf("Incorrect collision at position (%d,%d). Expected %t, got %t",
				location.gridX, location.gridY, location.isCollision, collision)
		}
	}
}

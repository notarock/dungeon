package dungeon

import (
	"math/rand"
)

type Tile int

const (
	Empty Tile = iota + 1
	Floor
	Wall
	Hallway
)

const (
	MIN_X = 10
	MAX_X = 45
	MIN_Y = 10
	MAX_Y = 45
)

func randomX() int {
	return (rand.Intn(MAX_X-MIN_X) + MIN_X)
}

func randomY() int {
	return (rand.Intn(MAX_Y-MIN_Y) + MIN_Y)
}

func GenerateEmptyTile(x, y int) [][]Tile {
	canvas := make([][]Tile, x)
	for i := range canvas {
		canvas[i] = make([]Tile, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			canvas[i][j] = Empty
		}
	}

	for i := 0; i < x; i++ {
		canvas[i][0] = Wall
		canvas[i][len(canvas[0])-1] = Wall
	}

	for i := 0; i < y; i++ {
		canvas[0][i] = Wall
		canvas[len(canvas)-1][i] = Wall
	}

	return canvas
}

package dungeon

import (
	"math/rand"
	"time"
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

var TILE_CHARSET = map[Tile]string{
	Empty:   " ",
	Floor:   "░",
	Wall:    "█",
	Hallway: "+",
}

type Level struct {
	tiles          [][]Tile
	RootPartitions BspNode
}

func randomX() int {
	return (rand.Intn(MAX_X-MIN_X) + MIN_X)
}

func randomY() int {
	return (rand.Intn(MAX_Y-MIN_Y) + MIN_Y)
}

func GenerateCanvas() [][]Tile {
	rand.Seed(time.Now().UnixNano())
	x := randomX()
	y := randomY()

	canvas := make([][]Tile, x)
	for i := range canvas {
		canvas[i] = make([]Tile, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			canvas[i][j] = Floor
		}
	}

	for i := 0; i < x; i++ {
		canvas[i][0] = Wall
		canvas[i][len(canvas[0])-1] = Wall
	}

	return canvas
}

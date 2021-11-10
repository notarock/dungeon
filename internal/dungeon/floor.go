package dungeon

import (
	// "fmt"
	"fmt"
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
	MIN_X = 20
	MAX_X = 90
	MIN_Y = 20
	MAX_Y = 90
)

func randomX() int {
	return rand.Intn(MAX_X - MIN_X) + MIN_X
}

func randomY() int {
	return rand.Intn(MAX_Y - MIN_Y) + MIN_Y
}

func generateCanvas() [][]Tile {
	x := randomX()
	y := randomY()

	var canvas [][]Tile

	for i := 0; i < x; i++ {
		row := make([]Tile, y)

		for j := 0; j < y; j++ {
			canvas[i][j] = Empty
		}
	}

	fmt.Printf("canvas: %v\n", canvas)

	return canvas
}

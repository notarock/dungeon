package game

import (
	"github.com/notarock/dungeon/pkg/dungeon"
)

const (
	TILE_PLAYER = "@"
)

var TILE_CHARSET = map[dungeon.Tile]string{
	dungeon.Empty:   " ",
	dungeon.Floor:   "░",
	dungeon.Wall:    "█",
	dungeon.Hallway: "+",
}

func DrawMap(floor [][]dungeon.Tile, p Player) string {
	var drawn string
	for x, row := range floor {
		for y, tile := range row {
			if x == p.GetX() && y == p.GetY() {
				drawn += "@"
			} else {
				drawn += TILE_CHARSET[tile]
			}
		}
		drawn += "\n"
	}

	return drawn
}

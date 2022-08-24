package dungeon

import "github.com/notarock/dungeon/pkg/dungeon/tile"

const (
	TILE_PLAYER = "@"
)

func DrawMap(floor [][]tile.Tile, p Player) string {
	var drawn string
	for x, row := range floor {
		for y, tile := range row {
			if x == p.GetX() && y == p.GetY() {
				drawn += "@"
			} else {
				drawn += tile.DrawTile()
			}
		}
		drawn += "\n"
	}

	return drawn
}

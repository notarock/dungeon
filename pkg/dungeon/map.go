package dungeon

import (
	"math"

	"github.com/notarock/dungeon/pkg/dungeon/tile"
)

const (
	TILE_PLAYER = "@"
)

// GenerateMap create a bunch of rooms and stick them together to create a dungeon
// Not implemented yet
type Map struct {
	Tiles          [][]tile.Tile
	StartX, StartY int
}

func (m *Map) GetTile(x, y int) (tile.Tile, error) {
	// if x > len(m.Tiles) || y > len(m.Tiles[0]) {
	// 	return tile.TileEmpty{}, fmt.Errorf("selected tile is out of bounds")
	// }

	lol := m.Tiles[x][y]
	return lol, nil
}

func drawMap(floor [][]tile.Tile, p Player) string {
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

func (m *Map) LightAroundPosition(tx, ty int) {
	for i := int(math.Max(0, float64(tx-1))); i < int(math.Min(float64(tx+2), float64(len(m.Tiles)))); i++ {
		for j := int(math.Max(0, float64(ty-1))); j < int(math.Min(float64(ty+2), float64(len(m.Tiles[0])))); j++ {
			t := m.Tiles[i][j]
			t.LightUp()
			m.Tiles[i][j] = t
		}
	}

}

package dungeon

import (
	"math"

	"github.com/notarock/dungeon/pkg/dungeon/tile"
)

const (
	TILE_PLAYER         = "@"
	PLAYER_VISION_RANGE = 6
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

func (m *Map) LightAroundPosition(tileX, tileY int) {
	for i := int(math.Max(0, float64(tileX-2))); i < int(math.Min(float64(tileX+PLAYER_VISION_RANGE), float64(len(m.Tiles)))); i++ {
		for j := int(math.Max(0, float64(tileY-2))); j < int(math.Min(float64(tileY+PLAYER_VISION_RANGE), float64(len(m.Tiles[0])))); j++ {
			m.LightTile(i, j)
		}
	}
}

func (m *Map) ClearAroundPosition(tileX, tileY int) {
	for i := int(math.Max(0, float64(tileX-2))); i < int(math.Min(float64(tileX+PLAYER_VISION_RANGE), float64(len(m.Tiles)))); i++ {
		for j := int(math.Max(0, float64(tileY-2))); j < int(math.Min(float64(tileY+PLAYER_VISION_RANGE), float64(len(m.Tiles[0])))); j++ {
			m.MarkTileAsDrawn(i, j)
		}
	}
}

func (m *Map) LightTile(i, j int) {
	tile := m.Tiles[i][j]
	tile.LightUp()
	m.Tiles[i][j] = tile
}

func (m *Map) MarkTileAsDrawn(i, j int) {
	tile := m.Tiles[i][j]
	tile.MarkDrawnTile()
	m.Tiles[i][j] = tile
}

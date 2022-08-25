package dungeon

import (
	"log"

	"github.com/notarock/dungeon/pkg/dungeon/tile"
)

const (
	mapx = 40
	mapy = 100
)

func NewMap() *Map {
	tiles := GenerateEmptyTile(mapx, mapy)
	m := Map{Tiles: tiles}
	b := NewBspNode(0, mapx, 0, mapy)

	b.Partition(5, 75)
	m.MakeRooms(b)

	sx, sy := b.GetStartingCoords()

	m.StartX = sx
	m.StartY = sy

	return &m
}

func (m *Map) MakeRooms(bsp BspNode) {
	log.Println("makeroom", bsp)

	if bsp.Front != nil {
		m.MakeRooms(*bsp.Front)
	}

	if bsp.Back != nil {
		m.MakeRooms(*bsp.Back)
	}

	m.MakeRoom(
		bsp.GetMinx(),
		bsp.GetMaxx(),
		bsp.GetMiny(),
		bsp.GetMaxy(),
	)
}

func (m *Map) MakeRoom(x1, x2, y1, y2 int) error {
	tiles := m.Tiles

	for i := x1; i < x2; i++ {
		tiles[i][y1] = tile.TileWall{}
		tiles[i][y2-1] = tile.TileWall{}
	}

	for i := y1; i < y2; i++ {
		tiles[x1][i] = tile.TileWall{}
		tiles[x2-1][i] = tile.TileWall{}
	}

	m.Tiles = tiles

	return nil
}

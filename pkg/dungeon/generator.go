package dungeon

import (
	"log"
)

// GenerateMap create a bunch of rooms and stick them together to create a dungeon
// Not implemented yet
type Map struct {
	Tiles [][]Tile
}

const (
	mapx = 40
	mapy = 100
)

func NewMap() (Map, error) {
	tiles := GenerateEmptyTile(mapx, mapy)
	m := Map{Tiles: tiles}
	b := NewBspNode(0, mapx, 0, mapy)

	b.Partition(5, 50)

	m.MakeRooms(b)

	return m, nil
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
		tiles[i][y1] = Wall
		tiles[i][y2-1] = Wall
	}

	for i := y1; i < y2; i++ {
		tiles[x1][i] = Wall
		tiles[x2-1][i] = Wall
	}

	m.Tiles = tiles

	return nil
}

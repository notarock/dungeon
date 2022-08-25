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

	m.joinRooms(b)

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
		tiles[i][y2-2] = tile.TileWall{}
	}

	for i := y1; i < y2; i++ {
		tiles[x1][i] = tile.TileWall{}
		tiles[x2-1][i] = tile.TileWall{}
		tiles[x2-2][i] = tile.TileWall{}
	}

	m.Tiles = tiles

	return nil
}

func (m *Map) joinRooms(bsp BspNode) {
	if bsp.Front == nil || bsp.Back == nil {
		return
	}

	sx := bsp.Front.minx + ((bsp.Front.maxx - bsp.Front.minx) / 2)
	sy := bsp.Front.miny + ((bsp.Front.maxy - bsp.Front.miny) / 2)

	dx := bsp.Back.minx + ((bsp.Back.maxx - bsp.Back.minx) / 2)
	dy := bsp.Back.miny + ((bsp.Back.maxy - bsp.Back.miny) / 2)

	switch {
	case sx == dx:
		if sy < dy {
			for i := 0; i < (dy - sy); i++ {
				m.Tiles[sx][sy+i] = tile.TileFloor{}
			}
		} else {
			for i := 0; i < (sy - dy); i++ {
				m.Tiles[sx][dy+i] = tile.TileFloor{}
			}
		}
	case sy == dy:
		if sy > dy {
			for i := 0; i < (dx - sx); i++ {
				m.Tiles[sx+i][sy] = tile.TileFloor{}
			}
		} else {
			for i := 0; i < (sx - dx); i++ {
				m.Tiles[dx+i][sy] = tile.TileFloor{}
			}
		}
	}

	// m.Tiles[sx][sy] = tile.TileWall{}
	// m.Tiles[dx][dy] = tile.TileWall{}

	m.joinRooms(*bsp.Front)
	m.joinRooms(*bsp.Back)

}

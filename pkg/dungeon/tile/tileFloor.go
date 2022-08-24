package tile

type TileFloor struct{}

const TILE_FLOOR = "FLOOR"

func (t TileFloor) GetType() string {
	return "FLOOR"
}

func (t TileFloor) Walkable() bool {
	return true
}

func (t TileFloor) DrawTile() string {
	return "░"
}

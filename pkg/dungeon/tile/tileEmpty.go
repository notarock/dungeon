package tile

type TileEmpty struct{}

const TILE_EMPTY = "EMPTY"

func (t TileEmpty) GetType() string {
	return "EMPTY"
}

func (t TileEmpty) Walkable() bool {
	return true
}

func (t TileEmpty) DrawTile() string {
	return " "
}

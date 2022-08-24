package tile

type TileWall struct{}

const TILE_WALL = "WALL"

func (t TileWall) GetType() string {
	return "WALL"
}

func (t TileWall) Walkable() bool {
	return false
}

func (t TileWall) DrawTile() string {
	return "█"
}

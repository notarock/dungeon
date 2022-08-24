package tile

type Tile interface {
	GetType() string
	Walkable() bool
	DrawTile() string
}

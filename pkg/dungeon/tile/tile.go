package tile

type Tile struct {
	tileType string
	lit      bool
	walkable bool
}

func (t Tile) GetType() string {
	return t.tileType
}

func (t Tile) IsWalkable() bool {
	return t.walkable
}

func (t Tile) DrawTile() string {
	if !t.Visible() {
		return " "
	}

	tile, _ := TILE_CHARACTERS_MAP[t.tileType]
	return tile
}

func (t Tile) Visible() bool {
	return t.lit
}

func (t *Tile) LightUp() {
	t.lit = true
}

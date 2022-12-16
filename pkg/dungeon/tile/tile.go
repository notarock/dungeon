package tile

const IN_RANGE_MODIFIER = "_IN_RANGE"

type Tile struct {
	tileType   string
	lit        bool
	wasJustLit bool
	walkable   bool
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

	tileName := t.tileType

	if t.wasJustLit {
		tile, ok := TILE_CHARACTERS_MAP[tileName+IN_RANGE_MODIFIER]
		if ok {
			return tile
		}
	}

	tile, _ := TILE_CHARACTERS_MAP[t.tileType]

	return tile
}

func (t Tile) Visible() bool {
	return t.lit
}

func (t *Tile) LightUp() {
	t.lit = true
	t.wasJustLit = true
}

func (t *Tile) MarkDrawnTile() {
	t.wasJustLit = false
}

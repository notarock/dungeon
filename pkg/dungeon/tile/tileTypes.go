package tile

var TILE_CHARACTERS_MAP = map[string]string{
	"EMPTY": " ",
	"WALL":  "█",
	"FLOOR": "░",
}

func NewEmptyTile() Tile {
	return Tile{
		tileType: "EMPTY",
		lit:      false,
		walkable: true,
	}
}

func NewFloorTile() Tile {
	return Tile{
		tileType: "FLOOR",
		lit:      false,
		walkable: true,
	}
}

func NewWallTile() Tile {
	return Tile{
		tileType: "WALL",
		lit:      false,
		walkable: false,
	}
}

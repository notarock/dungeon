package dungeon

import (
	"fmt"
)

type Game struct {
	level      int
	player     Player
	currentMap *Map
}

func (g *Game) GetMapSize() (x, y int) {
	gameX := len(g.currentMap.Tiles)
	gameY := len(g.currentMap.Tiles[0])
	return gameX, gameY
}

func NewGame() (Game, error) {
	cMap := NewMap()
	p, err := InitPlayer(cMap.StartX, cMap.StartY)

	if err != nil {
		return Game{}, err
	}

	newGame := Game{
		level:      0,
		player:     p,
		currentMap: cMap,
	}

	return newGame, nil
}

func (g *Game) Move(d moveDirection) error {
	var tx, ty int

	switch d {
	case Up:
		tx, ty = g.player.GetX()-1, g.player.GetY()
	case Down:
		tx, ty = g.player.GetX()+1, g.player.GetY()
	case Left:
		tx, ty = g.player.GetX(), g.player.GetY()-1
	case Right:
		tx, ty = g.player.GetX(), g.player.GetY()+1
	}

	tile, err := g.currentMap.GetTile(tx, ty)
	if err != nil {
		return fmt.Errorf("could not retreive destination game tile of player move: %v", err)
	}

	if tile.IsWalkable() {
		g.currentMap.ClearAroundPosition(g.player.GetX(), g.player.GetY())
		g.player.xPosition = tx
		g.player.yPosition = ty
		g.currentMap.LightAroundPosition(tx, ty)
	}

	return nil
}

func (g Game) DrawGame() string {
	output := drawMap(g.currentMap.Tiles, g.player)

	return output
}

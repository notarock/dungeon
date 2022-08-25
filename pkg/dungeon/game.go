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
	switch d {
	case Up:
		tile, err := g.currentMap.GetTile(g.player.GetX()-1, g.player.GetY())
		if err != nil {
			return fmt.Errorf("could not retreive destination game tile of player move: %v", err)
		}

		if tile.Walkable() {
			g.player.xPosition -= 1
		}
	case Down:
		tile, err := g.currentMap.GetTile(g.player.GetX()+1, g.player.GetY())
		if err != nil {
			return fmt.Errorf("could not retreive destination game tile of player move: %v", err)
		}

		if tile.Walkable() {
			g.player.xPosition += 1
		}
	case Left:
		tile, err := g.currentMap.GetTile(g.player.GetX(), g.player.GetY()-1)
		if err != nil {
			return fmt.Errorf("could not retreive destination game tile of player move: %v", err)
		}

		if tile.Walkable() {
			g.player.yPosition -= 1
		}
	case Right:
		tile, err := g.currentMap.GetTile(g.player.GetX(), g.player.GetY()+1)
		if err != nil {
			return fmt.Errorf("could not retreive destination game tile of player move: %v", err)
		}

		if tile.Walkable() {
			g.player.yPosition += 1
		}

	}

	return nil
}

func (g Game) DrawGame() string {
	return drawMap(g.currentMap.Tiles, g.player)
}

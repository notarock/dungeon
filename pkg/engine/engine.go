package engine

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/notarock/dungeon/pkg/dungeon"
	"github.com/notarock/dungeon/pkg/player"
	"log"
)

var TILE_CHARSET = map[dungeon.Tile]string{
	dungeon.Empty:   " ",
	dungeon.Floor:   "░",
	dungeon.Wall:    "█",
	dungeon.Hallway: "+",
}

var floor dungeon.Map
var gamePlayer player.Player

func InitGame() {
	floor, _ = dungeon.NewMap()

	gameX := len(floor.Tiles)
	gameY := len(floor.Tiles[0])
	p, err := player.InitPlayer(gameX/2, gameY/2)
	if err != nil {
		log.Panicln(err)
	}
	gamePlayer = p

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, moveLeft); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, moveRight); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func moveLeft(g *gocui.Gui, v *gocui.View) error {
	gamePlayer.Move(player.Left)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveRight(g *gocui.Gui, v *gocui.View) error {
	gamePlayer.Move(player.Right)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	gamePlayer.Move(player.Up)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	gamePlayer.Move(player.Down)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func layout(g *gocui.Gui) error {
	gameX := len(floor.Tiles)
	gameY := len(floor.Tiles[0])
	if v, err := g.SetView("game", 0, 0, gameY+1, gameX+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, drawMap(floor.Tiles, gamePlayer))
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func redrawMap(gv *gocui.View) {
	gv.Clear()
	gv.Write([]byte(drawMap(floor.Tiles, gamePlayer)))
}

func drawMap(floor [][]dungeon.Tile, p player.Player) string {
	var drawn string
	for x, row := range floor {
		for y, tile := range row {
			if x == p.GetX() && y == p.GetY() {
				drawn += "@"
			} else {
				drawn += TILE_CHARSET[tile]
			}
		}
		drawn += "\n"
	}

	return drawn
}

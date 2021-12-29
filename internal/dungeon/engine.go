package dungeon

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var floor [][]Tile
var player Player

func Main() {
	floor = GenerateCanvas()

	gameX := len(floor)
	gameY := len(floor[0])
	player = InitPlayer(gameX, gameY)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	gameX := len(floor)
	gameY := len(floor[0])
	if v, err := g.SetView("game", 0, 0, gameX, gameY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, drawMap(floor, player))
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func drawMap(floor [][]Tile, player Player) string {
	var drawn string
	for x, row := range floor {
		for y, tile := range row {
			if x == player.PosX && y == player.PosY {
				drawn += "@"
			} else {
				drawn += TILE_CHARSET[tile]
			}
		}
		drawn += "\n"
	}

	return drawn
}

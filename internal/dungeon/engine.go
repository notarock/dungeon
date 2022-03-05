package dungeon

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var floor [][]Tile
var player Player

func InitGame() {
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
	player.Move(floor, Left)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveRight(g *gocui.Gui, v *gocui.View) error {
	player.Move(floor, Right)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	player.Move(floor, Up)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	player.Move(floor, Down)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func layout(g *gocui.Gui) error {
	gameX := len(floor)
	gameY := len(floor[0])
	if v, err := g.SetView("game", 0, 0, gameY+1, gameX+1); err != nil {
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

func redrawMap(gv *gocui.View) {
	gv.Clear()
	gv.Write([]byte(drawMap(floor, player)))
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

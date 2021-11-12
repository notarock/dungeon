package main

import (
	"fmt"
	"github.com/notarock/dungeon/internal/dungeon"
)

func main() {
	room := dungeon.NewRoom(false, false, false, false)

	fmt.Println(room.ToString())
}

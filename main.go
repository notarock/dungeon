package main

import (
	"fmt"
	"github.com/notarock/dungeon/internal/dungeon"
)

func main() {
	var room1 = *dungeon.NewRoom(true,true,true,true)
	var room2 = *dungeon.NewRoom(false,false,false,false)

	fmt.Println(dungeon.ToString(room1))
	fmt.Println(dungeon.ToString(room2))
}

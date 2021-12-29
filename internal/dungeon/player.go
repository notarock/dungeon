package dungeon

type Player struct {
	PosX, PosY int
}

func InitPlayer(MapX, MapY int) Player {
	return Player{PosX: int(MapX / 2), PosY: int(MapY / 2)}
}

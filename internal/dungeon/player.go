package dungeon

type moveDirection int

const (
	Up moveDirection = iota + 1
	Down
	Left
	Right
)

type Player struct {
	PosX, PosY int
}

func InitPlayer(MapX, MapY int) Player {
	return Player{PosX: int(MapX / 2), PosY: int(MapY / 2)}
}

func (p *Player) Move(m [][]Tile, d moveDirection) {
	switch d {
	case Up:
		p.PosX -= 1
	case Down:
		p.PosX += 1
	case Left:
		p.PosY -= 1
	case Right:
		p.PosY += 1
	}
}

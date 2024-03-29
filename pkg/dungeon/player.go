package dungeon

import "fmt"

type moveDirection int

const (
	Up moveDirection = iota + 1
	Down
	Left
	Right
)

type Player struct {
	xPosition, yPosition int
}

func (p Player) GetX() int { return p.xPosition }
func (p Player) GetY() int { return p.yPosition }

func InitPlayer(x, y int) (Player, error) {
	var p Player
	if x < 0 || y < 0 {
		return p, fmt.Errorf("failed to create new player at [%d, %d]: position must be positive", x, y)
	}

	p.xPosition = x
	p.yPosition = y

	return p, nil
}

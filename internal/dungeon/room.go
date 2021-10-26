package dungeon

type Room struct {
	North bool
	South bool
	East, West bool
}

func toStringOpening(door bool) string {
	if !door {
		return " ";
	} else {
		return "@";
	}
}

func NewRoom(n, s, e, w bool) *Room {
	return &Room{
		North : n,
		South : s,
		East : e,
		West : w,
	}
}

// Ouaches
func ToString(room Room) string {
	return "@" + toStringOpening(room.North) + "@\n" + toStringOpening(room.East) + " " + toStringOpening(room.West) + "\n" + "@" + toStringOpening(room.South) + "@"
}

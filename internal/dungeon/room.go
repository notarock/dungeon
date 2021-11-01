package dungeon

// Room struct that contains directionnal openings.
type Room struct {
	North, South, East, West bool
}

// toStringOpening takes in a door bolean and responds with a character
// to paint either a wall or a door (opening)
func toStringOpening(door bool) string {
	if door {
		return " "
	}
	return "@"
}

// NewRoom creates a room struct out of directionnal openings.
// n,s,e,w are boolean thats indicate if the door in that direction is a door
// (Or a plain wall if false)
func NewRoom(n, s, e, w bool) *Room {
	return &Room{
		North: n,
		South: s,
		East:  e,
		West:  w,
	}
}

// ToString takes a room, and create a string reprentation for it.
func ToString(room Room) string {
	return "@" + toStringOpening(room.North) + "@\n" + toStringOpening(room.East) + " " + toStringOpening(room.West) + "\n" + "@" + toStringOpening(room.South) + "@"
}

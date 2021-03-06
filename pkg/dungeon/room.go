package dungeon

// Room struct that contains directionnal openings.
type opening bool
type Room struct {
	North, South, East, West opening
}

// toStringOpening takes in a door bolean and responds with a character
// to paint either a wall or a door (opening)
func (o opening) toString() string {
	if o {
		return " "
	}
	return "@"
}

// NewRoom creates a room struct out of directionnal openings.
// n,s,e,w are boolean thats indicate if the door in that direction is a door
// (Or a plain wall if false)
func NewRoom(n, s, e, w opening) *Room {
	return &Room{
		North: n,
		South: s,
		East:  e,
		West:  w,
	}
}

// ToString takes a room, and create a string reprentation for it.
func (r Room) ToString() string {
	return "@" + r.North.toString() + "@\n" + r.East.toString() + " " + r.West.toString() + "\n" + "@" + r.South.toString() + "@"
}

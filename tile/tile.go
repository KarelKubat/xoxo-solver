// Package tile represents one tile on the board.
package tile

// Tile is the type of tiles.
type Tile int

// Tile values may be Undefined (outside of the board), Empty (not taken spot), or Cross or Circle.
const (
	Undefined Tile = iota
	Empty
	Cross
	Circle
)

// String returns a tile value as a string.
func (t Tile) String() string {
	switch t {
	case Empty:
		return "_"
	case Cross:
		return "X"
	case Circle:
		return "O"
	}
	return "?"
}

// FromRune converts a rune to a tile value: _ or - mean empty, X or x mean cross, O or o mean circle.
// Other values mean undefined.
func FromRune(r rune) Tile {
	if r == '_' || r == '-' {
		return Empty
	}
	if r == 'X' || r == 'x' {
		return Cross
	}
	if r == 'O' || r == 'o' {
		return Circle
	}
	return Undefined
}

// Invert returns the opposite of a tile: for cross this is a circle and for a circle this is a
// cross. Other values don't have opposites.
func Invert(t Tile) Tile {
	switch t {
	case Cross:
		return Circle
	case Circle:
		return Cross
	}
	return t
}

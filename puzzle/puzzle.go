// Package puzzle represents the board and the tiles on it.
package puzzle

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/KarelKubat/xoxo-solver/tile"
)

// Puzzle is the receiver.
type Puzzle struct {
	Width, Height int
	Field         [][]tile.Tile
}

// NewFromFile reads a textfile with an initial configuration. The input file must be structured
// as follows:
// - There must be at least 2 rows of tiles.
// - The rows must have the same length of tiles.
// - Tiles are indicated by X, O, or _ (lowercase is also accepted, and - instead of _ is ok)
// The validity of the input is not checked.
func NewFromFile(f string) (*Puzzle, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(b), "\n")
	if len(rows) < 2 {
		return nil, errors.New("input needs to have at least 2 rows")
	}
	field := [][]tile.Tile{}
	for _, r := range rows {
		if r == "" {
			continue
		}
		frow := []tile.Tile{}
		for _, c := range r {
			if c == ' ' {
				continue
			}
			t := tile.FromRune(c)
			if t == tile.Undefined {
				return nil, fmt.Errorf("unsupported field specifier '%v' in the input", c)
			}
			frow = append(frow, t)
		}
		if len(field) != 0 && len(field[0]) != len(frow) {
			return nil, fmt.Errorf("attempt to add a row of %v values, but rows must be %v long", len(frow), len(field[0]))
		}
		field = append(field, frow)
	}

	return &Puzzle{
		Width:  len(field[0]),
		Height: len(field),
		Field:  field,
	}, nil
}

// String returns the board as one string. Tiles are indicated by X, O or _.
func (p *Puzzle) String() string {
	out := ""
	for r := 0; r < p.Height; r++ {
		if out != "" {
			out += "\n"
		}
		out += p.rowString(r)
	}
	return out
}

// rowString is a helper to gather one row of the board as a string.
// It's also used when checking that no two rows are identical (a-la a hash value).
func (p *Puzzle) rowString(r int) string {
	out := ""
	for _, t := range p.Field[r] {
		out += t.String()
		out += " "
	}
	return out
}

// colString is a helper to gather one column of the board as a string.
// It is used when checking that no two columns are identical.
func (p *Puzzle) colString(c int) string {
	out := ""
	for r := 0; r < p.Height; r++ {
		out += p.Field[r][c].String()
	}
	return out
}

// ValueAt returns the tile value at a given board position. Coordinates pointing outside of the
// board are returned as tile.Undefined.
func (p *Puzzle) ValueAt(row, col int) tile.Tile {
	if row < 0 || row >= p.Height || col < 0 || col >= p.Width {
		return tile.Undefined
	}
	return p.Field[row][col]
}

// HasValue returns true when a board position has either a cross or a circle tile. Empty tiles
// or coordinates pointing outside of the board are returned as false.
func (p *Puzzle) HasValue(row, col int) bool {
	v := p.ValueAt(row, col)
	return v == tile.Cross || v == tile.Circle
}

// SetValue tries to place a tile on the board and returns true when this succeeded. The following
// constraints must be met in order to place the tile:
// - No 3 X's or O's may occur horizontally or vertically
// - A full row (having only X's or O's) may not be repeated
// - A full column may not be repeated
func (p *Puzzle) SetValue(row, col int, t tile.Tile) bool {
	// No 3 X's or O's horizontally.
	if p.ValueAt(row, col-2) == t && p.ValueAt(row, col-1) == t ||
		p.ValueAt(row, col+1) == t && p.ValueAt(row, col+2) == t ||
		p.ValueAt(row, col-1) == t && p.ValueAt(row, col+1) == t {
		return false
	}
	// No 3 X's or O's vertically.
	if p.ValueAt(row-2, col) == t && p.ValueAt(row-1, col) == t ||
		p.ValueAt(row+1, col) == t && p.ValueAt(row+2, col) == t ||
		p.ValueAt(row-1, col) == t && p.ValueAt(row+1, col) == t {
		return false
	}

	// Assume that this is a valid move, take back later if needed.
	thisTile := p.Field[row][col]
	p.Field[row][col] = t

	// Rows may not be repeated. Only compare full rows.
	thisRowString := p.rowString(row)
	if !strings.Contains(thisRowString, tile.Empty.String()) {
		for r := 0; r < p.Height; r++ {
			if r == row {
				continue
			}
			if p.rowString(r) == thisRowString {
				p.Field[row][col] = thisTile
				return false
			}
		}
	}
	// Columns may not be repeated. Only contain full columns.
	thisColString := p.colString(col)
	if !strings.Contains(thisColString, tile.Empty.String()) {
		for c := 0; c < p.Width; c++ {
			if c == col {
				continue
			}
			if p.colString(c) == thisColString {
				p.Field[row][col] = thisTile
				return false
			}
		}
	}

	// All checks passed.
	return true
}

// ClearValue sets a board position to tile.Empty. It is used when taking back a move while solving.
func (p *Puzzle) ClearValue(row, col int) {
	p.Field[row][col] = tile.Empty
}

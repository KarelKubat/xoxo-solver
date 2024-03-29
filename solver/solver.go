// Package solver has the logic to solve the puzzle.
package solver

import (
	"github.com/KarelKubat/xoxo-solver/l"
	"github.com/KarelKubat/xoxo-solver/puzzle"
	"github.com/KarelKubat/xoxo-solver/tile"
)

// Solver is the receiver and contains a puzzle board to work on.
type Solver struct {
	puzzle *puzzle.Puzzle
}

// New initializes and returns a Solver.
func New(p *puzzle.Puzzle) *Solver {
	return &Solver{
		puzzle: p,
	}
}

// FillHorizontalMiddles populates middle tiles enclosed by X's or O's. E.g., --X-X-- can ever only
// lead to --XOX--. There must be an O in the middle position, since three consecutive same tiles
// are forbidden. It returns true when board changes were made.
func (p *Solver) FillHorizontalMiddles() bool {
	filled := false
	for r := 0; r < p.puzzle.Height; r++ {
		for c := 0; c < p.puzzle.Width; c++ {
			// Prefill only blank spots.
			if p.puzzle.HasValue(r, c) {
				continue
			}
			// Prefill horizontals: O-O can only become OXO
			if p.puzzle.HasValue(r, c-1) && p.puzzle.HasValue(r, c+1) &&
				p.puzzle.ValueAt(r, c-1) == p.puzzle.ValueAt(r, c+1) {
				inverted := tile.Invert(p.puzzle.ValueAt(r, c-1))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling horizontal middles: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
		}
	}
	return filled
}

// FillVerticalMiddles populates middle tiles, but vertically. It returns true when board changes
// were made.
func (p *Solver) FillVerticalMiddles() bool {
	filled := false
	for r := 0; r < p.puzzle.Height; r++ {
		for c := 0; c < p.puzzle.Width; c++ {
			// Prefill only blank spots.
			if p.puzzle.HasValue(r, c) {
				continue
			}
			// Prefill verticals: O-O (but vertically) can only become OXO.
			if p.puzzle.HasValue(r-1, c) && p.puzzle.HasValue(r+1, c) &&
				p.puzzle.ValueAt(r-1, c) == p.puzzle.ValueAt(r+1, c) {
				inverted := tile.Invert(p.puzzle.ValueAt(r-1, c))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling vertical middles: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
		}
	}
	return filled
}

// FillHorizontalSides populates prefixes or postfixes. E.g., --XX-- can only ever be solved
// by prefixing and postfixing with an O into -OXXO-, because 3 consecutive identical tiles are
// forbidden. It returns true when board changes were made.
func (p *Solver) FillHorizontalSides() bool {
	filled := false
	for r := 0; r < p.puzzle.Height; r++ {
		for c := 0; c < p.puzzle.Width; c++ {
			// Prefill only blank spots.
			if p.puzzle.HasValue(r, c) {
				continue
			}
			// Prefill horizontal prefix: -OO can only become XOO.
			if p.puzzle.HasValue(r, c+1) && p.puzzle.HasValue(r, c+2) &&
				p.puzzle.ValueAt(r, c+1) == p.puzzle.ValueAt(r, c+2) {
				inverted := tile.Invert(p.puzzle.ValueAt(r, c+1))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling horizontal sides: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
			// Prefill horizontal postfix: OO- can only become OOX.
			if p.puzzle.HasValue(r, c-2) && p.puzzle.HasValue(r, c-1) &&
				p.puzzle.ValueAt(r, c-2) == p.puzzle.ValueAt(r, c-1) {
				inverted := tile.Invert(p.puzzle.ValueAt(r, c-2))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling horizontal sides: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
		}
	}
	return filled
}

// FillVerticalSides populates prefixes or postfixes but vertically. It returns true when board
// changes were made.
func (p *Solver) FillVerticalSides() bool {
	filled := false
	for r := 0; r < p.puzzle.Height; r++ {
		for c := 0; c < p.puzzle.Width; c++ {
			// Prefill only blank spots.
			if p.puzzle.HasValue(r, c) {
				continue
			}
			// Prefill vertical prefix: -OO (but vertically) can only become XOO.
			if p.puzzle.HasValue(r-1, c) && p.puzzle.HasValue(r-2, c) &&
				p.puzzle.ValueAt(r-1, c) == p.puzzle.ValueAt(r-2, c) {
				inverted := tile.Invert(p.puzzle.ValueAt(r-1, c))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling vertical sides: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
			// Prefill vertical postfix: OO- (but vertically)can only become OOX.
			if p.puzzle.HasValue(r+1, c) && p.puzzle.HasValue(r+2, c) &&
				p.puzzle.ValueAt(r+1, c) == p.puzzle.ValueAt(r+2, c) {
				inverted := tile.Invert(p.puzzle.ValueAt(r+1, c))
				if p.puzzle.SetValue(r, c, inverted) {
					l.Printf("Filling vertical sides: %v added at [%v,%v]", inverted, r, c)
					filled = true
				}
			}
		}
	}
	return filled
}

// FillBlanks starts the recursive puzzle solver and returns true when a solution is found, else
// false.
func (p *Solver) FillBlanks() bool {
	// Start filling at the first blank spot.
	for r := 0; r < p.puzzle.Height; r++ {
		for c := 0; c < p.puzzle.Width; c++ {
			if !p.puzzle.HasValue(r, c) {
				l.Printf("Filling blanks starts at [%v,%v]", r, c)
				return p.fillAt(r, c, 0)
			}
		}
	}
	return true
}

// fillAt is a recursive helper. The arguments are the Y/X coordinates of a blank tile to start
// solving, and an iteration counter (nice for tracing). fillAt calls itself recursively to fill
// the "next" empty tile. The return value for each iteration is true when a tile could be placed,
// else false. Since fillAt calls itself, the overall result is true when the puzzle could be
// solved.
func (p *Solver) fillAt(row, col, iteration int) bool {
	// Beyond the board, or at already filled tiles, is the stop condition.
	if row >= p.puzzle.Height || p.puzzle.HasValue(row, col) {
		l.Printf("Iteration %v: done", iteration)
		return true
	}
	l.Printf("Iteration %v:\n%v", iteration, p.puzzle)

	// Find the next coords to fill.
	nextCol := col + 1
	nextRow := row
	for {
		if nextCol >= p.puzzle.Width {
			nextCol = 0
			nextRow++
		}
		if !p.puzzle.HasValue(nextRow, nextCol) {
			break
		}
		nextCol++
	}

	// Try to place valid tiles.
	for _, t := range []tile.Tile{tile.Cross, tile.Circle} {
		if !p.puzzle.SetValue(row, col, t) {
			continue
		}
		l.Printf("Iteration %v: placing %v at [%v,%v], next up will be [%v,%v]", iteration, t, row, col, nextRow, nextCol)
		if !p.fillAt(nextRow, nextCol, iteration+1) {
			l.Printf("Iteration %v: taking back %v from [%v,%v]", iteration, t, row, col)
			p.puzzle.ClearValue(row, col)
			continue
		}
		// Wow, this worked.
		return true
	}
	// Options are exhausted.
	return false
}

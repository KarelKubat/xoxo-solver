# XOXO Puzzle Solver

<!-- toc -->
- [Invoking the solver](#invoking-the-solver)
- [How the solver works](#how-the-solver-works)
- [The code](#the-code)
<!-- /toc -->

I came across this type of puzzle in April 2022. It's a variant of Sudoku, but just a tad different. I thought it would be neat to write up a solver. Plus it keeps me busy during a rainy weekend.

The puzzle goes as follows:

You get an initial board of of varying dimensions. The board width isn't necessarily the height; the board is just a rectangle. Some tiles on the board have an X or an O on it. Your job is to fill the entire board with X's or O's, so that:

- The repetition of X's or O's may not be longer than two, both horizontally and vertically. So you can't have 3 O's next to each other, or 3 X's above one another.
- Every row on the board must be unique. You can't have one row OOXXOO and another row OOXXOO.
- Every column on the board must be unique.

Here's an example of an 8x12 board (see also the file `sample1.txt`):

```
- - o - - - - -
- x - - - - - -
- - o - - o x -
o - - o - - x -
- - o o - - - -
- - - - - x - x
- - o - o - - x
- - - - - - o -
- x - - - x - -
- - - - - - - -
o - - - - - o -
- - - x - - o -
```

## Invoking the solver

The solver is invoked as follows:

- `go run xoxo-solver.go $INPUTFILE`. Alternatively you can of course `go build xoxo-solver.go` and invoke the binary.
- If you want to see some tracing, add the flag `--verbose`:
  - `go run -- xoxo-solver.go --verbose $INPUTFILE` needs the `--` as to not confuse `go run`
  - Or with a prebuilt binary, it's just `xoxo-solver --verbose $INPUTFILE`

The `$INPUTFILE` is just a text file with the initial board. See the files `*.txt`.

- X-tiles must be marked using `X` or `x`.
- O-tiles must be marked using `O` or `o`.
- Empty spots must be marked using `_` or `-`.
- You can put blanks in the between the tiles for readability.

## How the solver works

- The solver tries to "pre-populate" the board with obvious moves that you can't avoid. For example, XX_ (_ means empty) can only lead to XXO (you can't have three X's in a row). Same goes for verticals.
- Another obvious move is to change X_X into XOX, because again, you can't have three identical tiles in a row right behind one another (or in a column).
- Next, the solver fills the first blank ("first" arbitrarily meaning: from the top left). That places a tile on the board, and checks whether the next-up blank can be filled. If not, the tile is taken back and another tile variant (X instead of O, or v.v.) is tried.
- This last step is obviously a recursive invocation.

The solver will correctly find one solution, but it won't exhaustively find all solutions. It won't detect an incorrect initial configuration (e.g., one that contains OOO), but it will of course avoid creating incorrect boards while solving.

## The code

The Go code is fairly decent and modularized, but it doesn't have tests. The rainy weekend was over.
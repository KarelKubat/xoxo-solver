# XOXO Puzzle Solver

I came across this type of puzzle in April 2022. It's a variant of Sudoku, but just a tad different. I thought it would be neat to write up a solver.

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
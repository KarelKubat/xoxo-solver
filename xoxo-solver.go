package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/KarelKubat/xoxo-solver/l"
	"github.com/KarelKubat/xoxo-solver/puzzle"
	"github.com/KarelKubat/xoxo-solver/solver"
)

const usage = `
Usage:
  xoxo-solver [--verbose] INPUTFILE                # or
  go run -- xoxo-solver.go [--verbose] INPUTFILE
Flag --verbose enables tracing while solving.
`

func main() {
	var verboseFlag = flag.Bool("verbose", false, "enable tracing while solving")
	flag.Parse()
	l.Verbose(*verboseFlag)

	if flag.NArg() != 1 {
		log.Fatal(usage)
	}

	p, err := puzzle.NewFromFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	l.Printf("Initial configuration:\n%v", p)

	s := solver.New(p)
	for {
		filled := false
		if s.FillHorizontalMiddles() {
			filled = true
			l.Printf("Horizontal middles filled:\n%v", p)
		}
		if s.FillVerticalMiddles() {
			filled = true
			l.Printf("Vertical middles filled:\n%v", p)
		}
		if s.FillHorizontalSides() {
			filled = true
			l.Printf("Horizontal sides filled:\n%v", p)
		}
		if s.FillVerticalSides() {
			filled = true
			l.Printf("Vertical sides filled:\n%v", p)
		}
		if !filled {
			break
		}
	}

	if s.FillBlanks() {
		fmt.Printf("\n%v\n\n", p)
	} else {
		log.Fatal("No solution")
	}
}

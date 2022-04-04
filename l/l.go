// Package l wraps logging.
package l

import (
	"fmt"
	"log"
)

var verbose bool

// Verbose enables or disables logging by Printf().
func Verbose(v bool) {
	verbose = v
}

// Printf emits a-la log.Printf() but only if verbosity is turned on.
func Printf(f string, args ...interface{}) {
	if !verbose {
		return
	}
	log.Print(fmt.Sprintf(f, args...))
}

// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package runtime

import (
	"bytes"
	"log"
	"runtime/pprof"
)

func PrintStack() {
	var w bytes.Buffer
	p := pprof.Lookup("goroutine")
	err := p.WriteTo(&w, 1)
	if err != nil {
		log.Fatalf("Failed to PrintStack (%v)", err)
	}
	println(w.String())
}

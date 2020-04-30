// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	// "log"
	"strings"
	"unicode"

	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
)

//
type Filter struct{ be.Sparkless }

func (Filter) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(cir.Circuit)
	//
	name, view := x.NameAt("Name"), x.CircuitAt("View")
	nameStr, ok := name.(string)
	if !ok {
		return
	}
	if !strings.HasPrefix(nameStr, "Test") {
		return
	}
	sfx := nameStr[len("Test"):]
	if len(sfx) == 0 || !unicode.IsUpper(rune(sfx[0])) {
		return
	}
	y := cir.New().
		Grow("Address", x.CircuitAt("Address")).
		Grow("Name", name).
		Grow("View", view)
	eye.Show("Out", y)
}

func (Filter) CognizeOut(eye *be.Eye, v interface{}) {}

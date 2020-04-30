// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/hoijui/escher/pkg/a"
)

// DefaultValve is the name of the default valve
const DefaultValve = ""

// Verb is an interpretation of a circuit.
// The values of the sorted, number-named gates are treated as a slice of values representing an ‘address’.
// The value of the empty-string gate, if present, is expected to be a string and is a ‘verb’ word.
type Verb Circuit

// NewAddress returns a verb-view circuit (like an array in other languages)
// of the supplied names in the supplied order.
func NewAddress(addr ...Name) Verb {
	x := New()
	for i, n := range addr {
		x.Gate[i] = n
	}
	return Verb(x)
}

// NewVerbAddress returns a verb-view circuit (like an array in other languages)
// with the given name and the supplied names in the supplied order
func NewVerbAddress(verb string, addr ...Name) Verb {
	x := NewAddress(addr...)
	x.Gate[Super] = verb
	return x
}

// IsVerb returns true if the supplied value is a verb
func IsVerb(v Value) bool {
	u, ok := v.(Circuit)
	if !ok {
		return false
	}
	s, ok := u.StringOptionAt(Super)
	return ok && (s == "*" || s == "@")
}

func (verb Verb) Address() (addr []Name) {
	if Circuit(verb).IsNil() {
		return nil
	}
	for _, i := range Circuit(verb).SortedNumbers() {
		addr = append(addr, verb.Gate[i])
	}
	return
}

func (verb Verb) Verb() Value {
	return verb.Gate[Super]
}

func (verb Verb) compactible() bool {
	for _, n := range Circuit(verb).SortedNumbers() {
		s, ok := verb.Gate[n].(string)
		if !ok {
			return false
		}
		if strings.ContainsAny(s, "@*.\n") {
			return false
		}
	}
	return true
}

func (verb Verb) Print(w io.Writer, f Format) {
	if !verb.compactible() {
		Circuit(verb).Print(w, f)
		return
	}
	io.WriteString(w, verb.summarize())
}

func (verb Verb) String() string {
	if !verb.compactible() {
		return Circuit(verb).String()
	}
	return verb.summarize()
}

func (verb Verb) summarize() string {
	index := Circuit(verb).SortedNumbers()
	var w bytes.Buffer
	if v, ok := verb.Gate[Super]; ok {
		w.WriteString(fmt.Sprintf("%v", v))
	}
	for _, i := range index {
		x := verb.Gate[i]
		fmt.Fprintf(&w, "%v", x)
		if i+1 < len(index) {
			w.WriteString(a.RefineSymbolString)
		}
	}
	return w.String()
}

// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
	// . "github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/fs"
)

// Reflex is a bundle of not yet attached sense endpoints (synapses).
type Reflex map[Name]*Synapse

type Materializer interface {
	Materialize(*Matter) (Reflex, Value)
}

// Native represents a materializable object implemented as a Go type.
type Native interface {
	Spark(*Eye, *Matter, ...interface{}) Value // Initializer
}

// Matter describes the circuit context that commissioned the present materialization.
type Matter struct {
	Address Address // Address of the materialized design in memory
	Design interface{} // Design
	View Circuit // Valves connected to this design in the enclosing program
	Path []Name // Path to this reflex within the residual, recursively following gate names
	Super *Matter // Matter of the circuit that recalled this reflex as a peer
}

// Source returns an object describing the source location of the enclosing circuit, if available.
func (m *Matter) Source() Value {
	if m.Super == nil {
		return nil
	}
	u, ok := m.Super.Design.(Circuit)
	if !ok {
		return nil
	}
	s, ok := u.OptionAt(fs.Source{})
	if !ok {
		return nil
	}
	return s
}

func (m *Matter) String() string {
	r := New().
		Grow("Residue", pathCircuit(m.Path)).
		Grow("Memory", m.Address)
	if src := m.Source(); src != nil {
		r.Grow("Source", src)
	}
	return r.String()		
}

func (m *Matter) Circuit() Circuit {
	return New().
		Grow("Address", m.Address).
		Grow("Design", m.Design).
		Grow("View", m.View).
		Grow("Path", pathCircuit(m.Path))
}

func pathCircuit(p []Name) Circuit {
	r := New()
	for i, n := range p {
		r.Grow(i, n)
	}
	return r
}

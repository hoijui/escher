// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

// Materialize
type Materialize struct {
	lookup plumb.Client
}

func (m *Materialize) Spark(eye *be.Eye, _ *be.Matter, _ ...interface{}) Value {
	m.lookup.Init(
		func (v interface{}) {
			eye.Show("Lookup", v)
		},
	)
	return &Materialize{}
}

// Design: { * }
func (m *Materialize) CognizeDesign(eye *be.Eye, v interface{}) {
	reflex, residual := be.Materialize(lookupClient{&m.lookup}, v.(Circuit))
	eye.Show(DefaultValve, New().Grow("Reflex", reflex).Grow("Residual", residual))
}

func (m *Materialize) CognizeLookup(eye *be.Eye, v interface{}) {
	m.lookup.Cognize(v)
}

// In: ignored
// Out: { Reflex Reflex; Residual Circuit }
func (m *Materialize) Cognize(*be.Eye, interface{}) {}

type lookupClient struct {
	*plumb.Client
}

func (c lookupClient) Lookup(addr Address) Value {
	return c.Client.Fetch(addr)
}

// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
)

type System struct {
	barrier cir.Circuit
}

func (s *System) Spark(_ *be.Eye, matter cir.Circuit, _ ...interface{}) cir.Value {
	s.barrier = matter
	return nil
}

func (s *System) CognizeView(eye *be.Eye, value interface{}) {
	u := value.(cir.Circuit)
	residue := be.MaterializeSystem(u.At("Program"), u.CircuitAt("Index"), s.barrier)
	eye.Show("Residue", residue)
}

func (s *System) CognizeResidue(*be.Eye, interface{}) {}

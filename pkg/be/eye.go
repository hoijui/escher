// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/hoijui/escher/pkg/circuit"
)

// TODO Is this basically the implementation of a _Valve_? If so, it should probably be renamed, as the comment below makes sufficient reference to the Eye concept

// Eye is a runtime facility that delivers messages by invoking gate methods and
// provides methods that the gate can use to send messages out.
//
// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in the book:
//     Probably Approximately Correct -
//     Nature's Algorithms for Learning and Prospering in a Complex World
//     PDF: https://dl.acm.org/doi/book/10.5555/2536711
//     It is a book about some basic concepts for Machine Learning from 1984.
// The mind's eye is a synchronization device which sees changes as ordered
// and thus introduces the illusory perception of time (and, eventually, of the
// higher-level concepts of cause and effect).
type Eye struct {
	show map[cir.Name]nerve
}

type nerve chan *ReCognizer

type EyeCognizer func(eye *Eye, valve cir.Name, value interface{})

func NewEye(given Reflex) (eye *Eye) {
	eye = &Eye{
		show: make(map[cir.Name]nerve),
	}
	for vlv := range given {
		eye.show[vlv] = make(nerve, 1)
	}
	return
}

func (eye *Eye) Connect(given Reflex, cog EyeCognizer) {
	for vlvLoop, synLoop := range given {
		vlv, syn := vlvLoop, synLoop
		go func() {
			eye.show[vlv] <- syn.Connect(
				func(w interface{}) {
					cog(eye, vlv, w)
				},
			)
		}()
	}
}

func (eye *Eye) Show(valve cir.Name, v interface{}) {
	n := eye.show[valve]
	r := <-n
	defer func() {
		n <- r
	}()
	r.ReCognize(v)
}

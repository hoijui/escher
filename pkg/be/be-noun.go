// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	cir "github.com/hoijui/escher/pkg/circuit"
)

// Required matter: Noun
func materializeNoun(given Reflex, matter cir.Circuit) (residue interface{}) {
	noun := matter.At("Noun")
	for _, synLoop := range given {
		syn := synLoop
		go func() {
			syn.Connect(DontCognize).ReCognize(noun)
		}()
	}
	if len(given) > 0 {
		return nil
	}
	return noun
}

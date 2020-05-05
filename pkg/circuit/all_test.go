// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit_test

import (
	"fmt"
	"testing"

	cir "github.com/hoijui/escher/pkg/circuit"
)

func TestSame(t *testing.T) {
	if !cir.Same(cir.New().Grow("x", nil), cir.New().Grow("x", nil)) {
		t.Errorf("same")
	}
	if !cir.Same(cir.New().Grow("x", cir.DefaultValve), cir.New().Grow("x", cir.DefaultValve)) {
		t.Errorf("same")
	}
}

func TestVerb(t *testing.T) {
	a, b := cir.Circuit(cir.NewVerbAddress("@", "abc", "d", 1)), cir.Circuit(cir.NewVerbAddress("@", "abc", "d", 1))
	if !cir.Same(a, b) {
		t.Errorf("verb same")
	}
	fmt.Printf("%v\n", a)
}

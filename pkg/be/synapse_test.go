// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be_test

import (
	"testing"

	"github.com/hoijui/escher/pkg/be"
)

func TestSynapse(t *testing.T) {
	x, y := be.NewSynapse()
	p, q := be.NewSynapse()
	go be.Link(y, p)
	go be.Link(p, y)
	ch := make(chan int)
	go func() {
		x.Connect(be.DontCognize).ReCognize(1)
		ch <- 1
	}()
	go func() {
		q.Connect(be.DontCognize).ReCognize(1)
		ch <- 1
	}()
	<-ch
	<-ch
}

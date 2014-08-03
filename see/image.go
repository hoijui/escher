// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"strconv"
	"github.com/gocircuit/escher/star"
)

const MatchingName = "$"

func SeeImage(src *Src) (y *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			y = nil
		}
	}()
	x := star.Make()
	m := star.Make()
	x.Merge(MatchingName, m)
	t := src.Copy()
	t.Match("{")
	Space(t)
	var i int
	for {
		q := t.Copy()
		Space(q)
		name, peer, match := SeePeerOrMatching(q)
		if peer == nil && match == nil {
			break
		}
		Space(q)
		t.Become(q)
		if peer != nil {
			x.Merge(name, peer)
		} else {
			m.Merge(strconv.Itoa(i), match)
			i++
		}
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	if m.Len() == 1 { // no matchings
		star.Split(x, MatchingName)
	}
	return Wrap(ImagineStar(x)) // enclose image in a star singleton
}
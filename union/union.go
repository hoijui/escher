// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package union

// Name is one of: int or string
type Name interface{}

// Meaning is one of: string, int, float64, complex128, Union
type Meaning interface{}

// Super is a placeholder meaning for the super peer
type Super struct{}

func (Super) String() string {
	return "*"
}

// union ...
type union struct {
	peer map[Name]Meaning
	match map[Name]map[Name]Matching // peer -> valve -> opposing peer and valve
}

type Union struct {
	*union
}

// Matching ...
type Matching struct {
	Peer [2]Name
	Valve [2]Name
}

func (x Matching) Reverse() Matching {
	x.Peer[0], x.Peer[1] = x.Peer[1], x.Peer[0]
	x.Valve[0], x.Valve[1] = x.Valve[1], x.Valve[0]
	return x
}

// New ...
func New() Union {
	return Union{
		&union{
			peer: make(map[Name]Meaning),
			match: make(map[Name]map[Name]Matching),
		},
	}
}

var Nil Union // the nil union

func (u *union) IsNil() bool {
	return u == nil
}

func (u *union) IsEmpty() bool {
	return len(u.peer) == 0 && len(u.match) == 0
}

// Change adds a peer to this union.
func (u *union) Change(name Name, meaning Meaning) (before Meaning, overwrite bool) {
	before, overwrite = u.peer[name]
	u.peer[name] = meaning
	return
}

func (u *union) ChangeExclusive(name Name, meaning Meaning) {
	if _, over := u.Change(name, meaning); over {
		panic(1)
	}
}

// At ...
func (c *union) At(name Name) (Meaning, bool) {
	v, ok := c.peer[name]
	return v, ok
}

func (u *union) Forget(name Name) Meaning {
	forgotten := u.peer[name]
	delete(u.peer, name)
	return forgotten
}

// Match ...
func (c *union) Match(x Matching) {
	if x.Peer[0] == x.Peer[1] && x.Valve[0] == x.Valve[1] {
		panic("mismatch")
	}
	p := []map[Name]Matching{
		c.valves(x.Peer[0]), 
		c.valves(x.Peer[1]),
	}
	v := x.Valve
	if _, ok := p[0][v[0]]; ok {
		panic("dup")
	}
	if _, ok := p[1][v[1]]; ok {
		panic("dup")
	}
	p[0][v[0]], p[1][v[1]] = x, x.Reverse()
}

func (c *union) valves(p Name) map[Name]Matching {
	if _, ok := c.peer[p]; !ok {
		c.peer[p] = nil // placeholder
	}
	if _, ok := c.match[p]; !ok {
		c.match[p] = make(map[Name]Matching)
	}
	return c.match[p]
}

func (u *union) Valves(peer Name) map[Name]Matching {
	return u.match[peer]
}

// Follow ...
func (c *union) Follow(p, v Name) (q, u Name) {
	x, ok := c.valves(p)[v]
	if !ok {
		return nil, nil
	}
	return x.Peer[1], x.Valve[1]
}

func (c *union) Letters() []string {
	var l []string
	for key, _ := range c.peer {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (c *union) Numbers() []int {
	var l []int
	for key, _ := range c.peer {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u *union) Symbols() map[Name]Meaning {
	return u.peer
}

func (u *union) String() string {
	return u.Print(nil, "", "\t")
}
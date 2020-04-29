// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

// Orient is a registry for gates names and their valve names
type Orient map[Name]map[Name]struct{} // gate -> valve -> yes/no?

// Include registers the gate supplied by name as having the supplied valve name
func (o Orient) Include(gate, valve Name) {
	valves, ok := o[gate]
	if !ok {
		valves = make(map[Name]struct{})
		o[gate] = valves
	}
	valves[valve] = struct{}{}
}

// Has returns true if the gate supplied by name has a valve with the supplied name
func (o Orient) Has(gate, valve Name) bool {
	if valves, ok := o[gate]; ok {
		if _, ok := valves[valve]; ok {
			return true
		}
	}
	return false
}

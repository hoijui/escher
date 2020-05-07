// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package circuit provides Escher gates
// for building dynamic cloud applications
// using the circuit runtime of http://gocircuit.org

// The following excludes this file from defualt compilation.
// To include, you would have to use this command:
// go build -tags=plugin_faculty_gocircuit ./...
// But We actually compile this into a Go plugin with:
// go build -v -buildmode=plugin -o bin/plugins/faculty/gocircuit.so ./pkg/faculty/gocircuit/
// +build plugin_faculty_gocircuit

package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hoijui/circuit/client"
	"github.com/hoijui/escher/pkg/be"
	"github.com/hoijui/escher/pkg/faculty"
)

// client *client.Client
func Init(discover []string) {
	program = &Program{}
	if len(discover) > 0 && discover[0] != "" {
		program.Client = client.DialDiscover(discover[0], nil)
	}

	rand.Seed(time.Now().UnixNano())

	faculty.Register(be.NewMaterializer(&Process{}), "element", "Process")
	faculty.Register(be.NewMaterializer(&Docker{}), "element", "Docker")
}

// Program…
type Program struct {
	*client.Client
}

var program *Program

// ChooseID returns a unique textual ID.
func ChooseID() string {
	return strconv.FormatUint(uint64(rand.Int63()), 20)
}

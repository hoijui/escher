// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
	"github.com/hoijui/escher/pkg/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&WriteFile{}), "io", "WriteFile")
}

type WriteFile struct {
	name  string
	named chan struct{}
}

func (h *WriteFile) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	h.named = make(chan struct{})
	return &WriteFile{}
}

func (h *WriteFile) CognizeName(eye *be.Eye, v interface{}) {
	h.name = v.(string)
	close(h.named)
}

func (h *WriteFile) CognizeContent(eye *be.Eye, v interface{}) {
	<-h.named
	var err error
	switch t := v.(type) {
	case string:
		err = ioutil.WriteFile(h.name, []byte(t), 0644)
	case []byte:
		err = ioutil.WriteFile(h.name, t, 0644)
	case io.Reader:
		var w bytes.Buffer
		io.Copy(&w, t)
		err = ioutil.WriteFile(h.name, w.Bytes(), 0644)
	default:
		panic("eh?")
	}
	if err != nil {
		log.Fatal("Failed to cognize content", err)
	}
	eye.Show("Ready", 1)
}

func (h *WriteFile) CognizeReady(eye *be.Eye, v interface{}) {}

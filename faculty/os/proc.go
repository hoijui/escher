// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	kitio "github.com/gocircuit/escher/kit/io"
	"github.com/gocircuit/escher/kit/plumb"
)

// Process
type Process struct{}

func (x Process) Materialize(*be.Matter) be.Reflex {
	p := &process{
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Command", "When", "Exit", "IO")
	return reflex
}

type process struct{
	spawn chan interface{}
	sync.Mutex
	cmd *exec.Cmd
}

func (p *process) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch valve {
	case "Command":
		if p.cognizeCommand(value) {
			p.startBack(eye)
		}
	case "When":
		p.spawn <- value
		// log.Printf("os process spawning (%v)", Linearize(fmt.Sprintf("%v", value)))
	}
}

//	{
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
func (p *process) cognizeCommand(v interface{}) (ready bool) {
	img, ok := v.(Image)
	if !ok {
		panic(fmt.Sprintf("Non-image sent to Process.Command (%v)", v))
		return false
	}
	p.Lock()
	defer p.Unlock()
	if p.cmd != nil {
		panic("process command already set")
	}
	p.cmd = &exec.Cmd{}
	p.cmd.Path = img.String("Path") // mandatory
	p.cmd.Args = []string{p.cmd.Path}
	if img.Has("Dir") {
		p.cmd.Dir = img.String("Dir")
	}
	env := img.Walk("Env")
	for _, key := range env.Numbers() {
		p.cmd.Env = append(p.cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Numbers() {
		p.cmd.Args = append(p.cmd.Args, args.String(key))
	}
	// log.Printf("os process command (%v)", Linearize(img.Print("", "")))
	return true
}

func (p *process) startBack(eye *plumb.Eye) {
	p.Lock()
	defer p.Unlock()
	f := &processFixed{
		eye: eye,
		cmd: p.cmd,
	}
	go f.backLoop(p.spawn)
}

type processFixed struct {
	eye *plumb.Eye
	cmd *exec.Cmd
}

func (p *processFixed) backLoop(spawn <-chan interface{}) {
	for {
		when := <-spawn
		var x Image
		if exit := p.spawnProcess(when); exit != nil {
			x = Image{
				"When": when,
				"Exit":  1,
			}
			p.eye.Show("Exit", x)
		} else {
			x = Image{
				"When": when,
				"Exit":  0,
			}
			p.eye.Show("Exit", x)
		}
		// log.Printf("os process exit sent (%v)", Linearize(x.Print("", "")))
	}
}

func (p *processFixed) spawnProcess(when interface{}) (err error) {
	var stdin io.WriteCloser
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	stdin, err =  p.cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err = p.cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err = p.cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err = p.cmd.Start(); err != nil {
		log.Fatalf("Problem starting %s (%v)", p.cmd.Path, err)
		return err
	}
	// We cannot call cmd.Wait before all std streams have been closed.
	stdClose := make(chan struct{}, 3)
	stdin = kitio.RunOnCloseWriter(stdin, func() { stdClose <- struct{}{} })
	stdout = kitio.RunOnCloseReader(stdout, func() { stdClose <- struct{}{} })
	stderr = kitio.RunOnCloseReader(stderr, func() { stdClose <- struct{}{} })
	g := Image{
		"When":  when,
		"Stdin":  stdin,
		"Stdout": stdout,
		"Stderr": stderr,
	}
	// log.Printf("os process io (%v)", Linearize(fmt.Sprintf("%v", when)))
	p.eye.Show("IO", g)
	<-stdClose
	<-stdClose
	<-stdClose
	// log.Printf("os process waiting (%v)", Linearize(fmt.Sprintf("%v", when)))
	err = p.cmd.Wait()
	switch err.(type) {
	case nil, *exec.ExitError:
	default:
		panic(err)
	}
	// log.Printf("os process exit (%v)", Linearize(fmt.Sprintf("%v", when)))
	return err
}

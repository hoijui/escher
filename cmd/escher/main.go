// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"plugin"
	"regexp"
	"strings"

	"github.com/hoijui/escher/pkg/a"
	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
	fac "github.com/hoijui/escher/pkg/faculty"
	"github.com/hoijui/escher/pkg/kit/fs"
	kio "github.com/hoijui/escher/pkg/kit/io"
	"github.com/hoijui/escher/pkg/see"

	// Load faculties
	_ "github.com/hoijui/escher/pkg/faculty/basic"
	_ "github.com/hoijui/escher/pkg/faculty/cmplx"
	_ "github.com/hoijui/escher/pkg/faculty/escher"
	"github.com/hoijui/escher/pkg/faculty/gocircuit"
	_ "github.com/hoijui/escher/pkg/faculty/http"
	_ "github.com/hoijui/escher/pkg/faculty/index"
	_ "github.com/hoijui/escher/pkg/faculty/io"
	_ "github.com/hoijui/escher/pkg/faculty/math"
	_ "github.com/hoijui/escher/pkg/faculty/model"
	fos "github.com/hoijui/escher/pkg/faculty/os"
	"github.com/hoijui/escher/pkg/faculty/test"
	_ "github.com/hoijui/escher/pkg/faculty/text"
	_ "github.com/hoijui/escher/pkg/faculty/time"
	_ "github.com/hoijui/escher/pkg/faculty/yield"
)

func main() {
	flagSrc := flag.String("src", "", "source directory")
	flagDiscover := flag.String("d", "", "multicast UDP discovery address for gocircuit.org faculty")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %v [OPTION]... <main-circuit> [ARGUMENT]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -h, --help\n\tprint this usage help message and exit\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  # When in escher sources repo root:\n")
		fmt.Fprintf(os.Stderr, "  %v -src ./src/ \"*tutorial.HelloWorldMain\"\n", os.Args[0])
	}
	flag.Parse()
	var flagMain string
	var flagArgs = flag.Args()
	if len(flagArgs) > 0 {
		flagMain = flagArgs[0]
		flagArgs = flagArgs[1:]
	}
	// parse env
	}
	// run main
	if main != "" {
		verb := see.ParseVerb(main)
		if cir.Circuit(verb).IsNil() {
			fmt.Fprintf(os.Stderr, "verb '%v' not recognized\n", verb)
			os.Exit(1)
		}
		exec(index, cir.Circuit(verb), false)
	}
	// standard loop
	r := kio.NewChunkReader(os.Stdin)
	cont := true
	for cont {
		chunk, err := r.Read()
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				// This is normal behavior; simply denotes the end of input
				cont = false
			} else {
				fatalf("end of session (%v)\n", err)
			}
		}
		src := a.NewSrcString(string(chunk))
		for src.Len() > 0 {
			u := see.SeeChamber(src)
			if u == nil || u.(cir.Circuit).Len() == 0 {
				break
			}
			fmt.Fprintf(os.Stderr, "MATERIALIZING %v\n", u)
			exec(index, u.(cir.Circuit), true)
		}
	}
}

func exec(index be.Index, verb cir.Circuit, showResidue bool) {
	residue := be.MaterializeSystem(cir.Circuit(verb), cir.Circuit(index), cir.New().Grow("Main", cir.New()))
	if showResidue {
		fmt.Fprintf(os.Stderr, "RESIDUE %v\n\n", residue)
	}
}

package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface capable of tracing application workflow events
type Tracer interface {
	Trace(...interface{})
}

// New is a helper function to create a tracer object
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a tracer that does nothing when the Trace is called
func Off() Tracer {
	return &nilTracer{}
}

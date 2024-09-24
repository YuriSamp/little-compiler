package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basepointer int) *Frame {
	return &Frame{cl: cl, ip: -1, basePointer: basepointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}

package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

type Definition struct {
	Name           string
	OpearandWidths []int
}

const (
	OpConstant Opcode = iota
	OpAdd
	OpPop
	OpSub
	OpMul
	OpDiv
	OpTrue
	OpFalse
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpMinus
	OpBang
	OpJumpNotTruthy
	OpJump
	OpNull
	OpGetGlobal
	OpSetGlobal
)

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpAdd: {"OpAdd", []int{}},
	OpPop: {"OpPop", []int{}},
	OpSub: {"OpSub", []int{}},
	OpDiv: {"OpDiv", []int{}},
	OpMul: {"OpMul", []int{}},
	OpTrue : {"OpTrue", []int{}},
	OpFalse : {"OpFalse", []int{}},
	OpEqual : {"OpEqual", []int{}},
	OpNotEqual : {"OpNotEqual", []int{}},
	OpGreaterThan : {"OpGreaterThan", []int{}},
	OpMinus: {"OpMinus", []int{}},
	OpBang: {"OpBang", []int{}},
	OpJumpNotTruthy  : {"OpJumpNotTruthy", []int{2}},
	OpJump: {"OpJump", []int{2}},
	OpNull : {"OpNull", []int{}},
	OpGetGlobal : {"OpGetGlobal", []int{2}},
	OpSetGlobal : {"OpSetGlobal", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]

	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]

	if !ok {
		return []byte{}
	}

	InstructionsLen := 1

	for _, w := range def.OpearandWidths {
		InstructionsLen += w
	}

	instuctions := make([]byte, InstructionsLen)

	instuctions[0] = byte(op)

	offset := 1

	for i, o := range operands {
		width := def.OpearandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instuctions[offset:], uint16(o))
		}
		offset += width
	}
	return instuctions
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int){
	operands := make([]int, len(def.OpearandWidths))

	offset :=0

	for i, width := range def.OpearandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		
		offset += width
	}

	return operands, offset
}

func ReadUint16 (ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0

	for i < len(ins){
		def, err := Lookup(ins[i])

		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n",err)
			continue
		}
		
		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n",i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandsCount := len(def.OpearandWidths)

	if len(operands) != operandsCount {
		return fmt.Sprintf("ERROR: operand len %d not match defined %d\n", len(operands), operandsCount)
	}

	switch operandsCount {
	case 0:
		return def.Name
	case 1 :
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

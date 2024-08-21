package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
	}

	for _, tt := range tests {
		instructions := Make(tt.op, tt.operands...)

		if len(instructions) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d", len(tt.expected), len(instructions))
		}

		for i, b := range tt.expected {
			if instructions[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d got=%d", i, b, instructions[i])
			}
		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpAdd),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}

	expected := `0000 OpAdd
0001 OpConstant 2
0004 OpConstant 65535
`

	concatted := Instructions{}

	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formated.\nwant=%q\ngot=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		byteRead int
	}{
		{OpConstant, []int{65535}, 2},
	}
	
	for _, tt := range tests{
		instructions := Make(tt.op, tt.operands...)
		
		def, err :=Lookup(byte(tt.op))
		
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}
		
		operandsRead, n := ReadOperands(def, instructions[1:])
		
		if n != tt.byteRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.byteRead, n)
		}
		
		for i, want :=range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operand wrong. want=%d, got=%d", want, operandsRead[i])
			}
		}
	}
}
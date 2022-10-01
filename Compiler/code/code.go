package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OpPop
	OpAdd
	OpSub
	OpMul
	OpDiv
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{4}},
	OpPop:      {"OpPop", []int{}},
	OpAdd:      {"OpAdd", []int{}},
	OpSub:      {"OpSub", []int{}},
	OpMul:      {"OpMul", []int{}},
	OpDiv:      {"OpDiv", []int{}},
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

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 4:
			binary.BigEndian.PutUint32(instruction[offset:], uint32(o))
		}
		offset += width
	}

	return instruction
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 4:
			operands[i] = int(ReadUint32(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint32(ins Instructions) uint32 {
	return binary.BigEndian.Uint32(ins)
}

func (ins Instructions) String() string {
	return ""
}

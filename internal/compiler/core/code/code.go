package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

/*
OpConstant
OpPop
OpAdd         +
OpSub         -
OpMul         *
OpSub         /
OpTrue        true
OpFalse       false
OpEqual       ==
OpNotEqual    !=
OpMinus       -
OpBamg        !
*/
const (
	OpConstant Opcode = iota
	OpPop
	OpJumpNotTruthy
	OpJump
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpRem
	OpTrue
	OpFalse
	OpNull
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpGreaterThanOrEquel
	OpMinus
	OpBang
	OpGetGlobal
	OpSetGlobal
	OpArray
	OpIndex
	OpCall
	OpReturn
	OpReturnValue
	OpGetLocal
	OpSetLocal
	OpGetBuiltin
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:           {"OpConstant", []int{4}},
	OpPop:                {"OpPop", []int{}},
	OpJumpNotTruthy:      {"OpJumpNotTruthy", []int{2}},
	OpJump:               {"OpJump", []int{2}},
	OpAdd:                {"OpAdd", []int{}},
	OpSub:                {"OpSub", []int{}},
	OpMul:                {"OpMul", []int{}},
	OpDiv:                {"OpDiv", []int{}},
	OpRem:                {"OpRem", []int{}},
	OpTrue:               {"OpTrue", []int{}},
	OpFalse:              {"OpFalse", []int{}},
	OpNull:               {"OpNull", []int{}},
	OpEqual:              {"OpEqual", []int{}},
	OpNotEqual:           {"OpNotEqual", []int{}},
	OpGreaterThan:        {"OpGreaterThan", []int{}},
	OpGreaterThanOrEquel: {"OpGreaterThanOrEquel", []int{}},
	OpMinus:              {"OpMinus", []int{}},
	OpBang:               {"OpBang", []int{}},
	OpGetGlobal:          {"OpGetGlobal", []int{4}},
	OpSetGlobal:          {"OpSetGlobal", []int{4}},
	OpArray:              {"OpArray", []int{4}},
	OpIndex:              {"OpIndex", []int{}},
	OpCall:               {"OpCall", []int{2}},
	OpReturn:             {"OpReturn", []int{}},
	OpReturnValue:        {"OpReturnValue", []int{}},
	OpGetLocal:           {"OpGetLocal", []int{2}},
	OpSetLocal:           {"OpSetLocal", []int{2}},
	OpGetBuiltin:         {"OpGetBuiltin", []int{2}},
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
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
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

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func ReadUint32(ins Instructions) uint32 {
	return binary.BigEndian.Uint32(ins)
}

package main

import (
	"fmt"
)

// InstructionOpcode encapsulates the opcode of a binary instruction
type InstructionOpcode uint8

func makeInstructionOpcode(s string) (InstructionOpcode, error) {
	i, err := HexExtract8(s)
	if err != nil {
		return InstructionOpcode(0), err
	}
	return InstructionOpcode(BitExtract8(i, 0x3f)), nil
}

// InstructionFactory describes a function that builds an instruction
// for the given binary components
type InstructionFactory func(InstructionArgs) InstructionBase

// UnrecognizedInstructionOpcodeError indicates the opcode was not valid
type UnrecognizedInstructionOpcodeError struct {
	op InstructionOpcode
}

func (err UnrecognizedInstructionOpcodeError) Error() string {
	return fmt.Sprintf("The opcode %v is unrecognized", err.op)
}

// Factory returns the InstructionFactory for the given opcode
func (op InstructionOpcode) Factory() (InstructionFactory, error) {
	switch op {
	case 0x00:
		return makeInstructionRD, nil
	case 0x01:
		return makeInstructionWR, nil
	case 0x02:
		return makeInstructionST, nil
	case 0x03:
		return makeInstructionLW, nil
	case 0x04:
		return makeInstructionMOV, nil
	case 0x05:
		return makeInstructionADD, nil
	case 0x06:
		return makeInstructionSUB, nil
	case 0x07:
		return makeInstructionMUL, nil
	case 0x08:
		return makeInstructionDIV, nil
	case 0x09:
		return makeInstructionAND, nil
	case 0x0a:
		return makeInstructionOR, nil
	case 0x0b:
		return makeInstructionMOVI, nil
	case 0x0c:
		return makeInstructionADDI, nil
	case 0x0d:
		return makeInstructionMULI, nil
	case 0x0e:
		return makeInstructionDIVI, nil
	case 0x0f:
		return makeInstructionLDI, nil
	case 0x10:
		return makeInstructionSLT, nil
	case 0x11:
		return makeInstructionSLTI, nil
	case 0x12:
		return makeInstructionHLT, nil
	case 0x13:
		return makeInstructionNOP, nil
	case 0x14:
		return makeInstructionJMP, nil
	case 0x15:
		return makeInstructionBEQ, nil
	case 0x16:
		return makeInstructionBNE, nil
	case 0x17:
		return makeInstructionBEZ, nil
	case 0x18:
		return makeInstructionBNZ, nil
	case 0x19:
		return makeInstructionBGZ, nil
	case 0x1a:
		return makeInstructionBLZ, nil
	default:
		return nil, UnrecognizedInstructionOpcodeError{op: op}
	}
}

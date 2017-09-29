package main

import (
	"regexp"
)

// InstructionComponents encapsulates the structure of a binary command.
type InstructionComponents struct {
	opcode InstructionOpcode
	args   InstructionArgs
}

func makeInstructionComponents(cmd string) (InstructionComponents, error) {
	re := regexp.MustCompile("^(?:0x)([0-9a-fA-F]{2})([0-9a-fA-F]{6})$")
	matches := re.FindStringSubmatch(cmd)

	opcode, opErr := makeInstructionOpcode(matches[1])
	if opErr != nil {
		return InstructionComponents{}, opErr
	}

	args, argErr := makeInstructionArgs(matches[2])
	if argErr != nil {
		return InstructionComponents{}, argErr
	}

	return InstructionComponents{opcode: opcode, args: args}, nil
}

func (ic InstructionComponents) decode() (InstructionBase, error) {
	factory, err := ic.opcode.Factory()
	if err != nil {
		return nil, err
	}
	return factory(ic.args), nil
}

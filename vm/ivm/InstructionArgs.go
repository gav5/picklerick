package ivm

import (
	"../../util"
)

// InstructionArgs encapsulates the args given to a binary instruction
type InstructionArgs uint32

func (args InstructionArgs) extractRegister(mask uint32) RegisterDesignation {
	return RegisterDesignation(util.BitExtract32(uint32(args), mask))
}

func (args InstructionArgs) extractAddress(mask uint32) Address {
	return Address(util.BitExtract32(uint32(args), mask))
}

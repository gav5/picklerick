package ivm

import (
	"fmt"
)

// InstructionArgsIO encapsulates the args for IO instructions
type InstructionArgsIO struct {
	Register1 RegisterDesignation
	Register2 RegisterDesignation
	Address   Address
}

// IOFormat returns the args in an IO format
func (args InstructionArgs) IOFormat() InstructionArgsIO {
	return InstructionArgsIO{
		Register1: args.extractRegister(0xf00000),
		Register2: args.extractRegister(0x0f0000),
		Address:   args.extractAddress(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsIO) ASM() string {
	if args.Register2 == 0x0 {
		return fmt.Sprintf(
			"%s %s",
			args.Register1.ASM(),
			args.Address.Dec(),
		)
	}
	return fmt.Sprintf(
		"%s (%s)",
		args.Register1.ASM(),
		args.Register2.ASM(),
	)
}

// ASMHex returns the representation in assembly language with hex args.
func (args InstructionArgsIO) ASMHex() string {
	if args.Register2 == 0x0 {
		return fmt.Sprintf(
			"%s %s",
			args.Register1.ASM(),
			args.Address.Hex(),
		)
	}
	return fmt.Sprintf(
		"%s (%s)",
		args.Register1.ASM(),
		args.Register2.ASM(),
	)
}

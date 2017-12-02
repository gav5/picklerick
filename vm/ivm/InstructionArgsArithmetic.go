package ivm

import (
	"fmt"
)

// InstructionArgsArithmetic encapsulates the args for arithmetic instructions
type InstructionArgsArithmetic struct {
	Source1     RegisterDesignation
	Source2     RegisterDesignation
	Destination RegisterDesignation
}

// ArithmeticFormat returns the args in an arithmetic format
func (args InstructionArgs) ArithmeticFormat() InstructionArgsArithmetic {
	return InstructionArgsArithmetic{
		Source1:     args.extractRegister(0xf00000),
		Source2:     args.extractRegister(0x0f0000),
		Destination: args.extractRegister(0xf0f000),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsArithmetic) ASM() string {
	return fmt.Sprintf(
		"%s %s %s",
		args.Destination.ASM(),
		args.Source1.ASM(),
		args.Source2.ASM(),
	)
}

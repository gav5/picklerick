package ivm

import (
	"fmt"
)

// InstructionArgsBranch encapsulates the args for branch instructions
type InstructionArgsBranch struct {
	Base        RegisterDesignation
	Destination RegisterDesignation
	Address     Address
}

// BranchFormat returns the args in a branch format
func (args InstructionArgs) BranchFormat() InstructionArgsBranch {
	return InstructionArgsBranch{
		Base:        args.extractRegister(0xf00000),
		Destination: args.extractRegister(0x0f0000),
		Address:     args.extractAddress(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsBranch) ASM() string {
	if args.Base == 0x0 {
		return fmt.Sprintf(
			"%s %s",
			args.Destination.ASM(),
			args.Address.Dec(),
		)
	}
	return fmt.Sprintf(
		"%s (%s)",
		args.Destination.ASM(),
		args.Base.ASM(),
	)
}

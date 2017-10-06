package instrType

import (
	"fmt"

	"../bus"
	"../reg"
)

// ArgsIO encapsulates the args for IO instructions
type ArgsIO struct {
	Register1 reg.Designation
	Register2 reg.Designation
	Address   bus.Address
}

// IOFormat returns the args in an IO format
func (args Args) IOFormat() ArgsIO {
	return ArgsIO{
		Register1: args.registerExtract(0xf00000),
		Register2: args.registerExtract(0x0f0000),
		Address:   args.addressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsIO) ASM() string {
	if args.Register2 == 0x0 {
		return fmt.Sprintf("%s %s", args.Register1.ASM(), args.Address.Hex())
	}
	return fmt.Sprintf("%s (%s)", args.Register1.ASM(), args.Register2.ASM())
}

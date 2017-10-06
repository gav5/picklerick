package instrType

import (
	"fmt"

	"../reg"
)

// ArgsArithmetic encapsulates the args for arithmetic instructions
type ArgsArithmetic struct {
	Source1     reg.Designation
	Source2     reg.Designation
	Destination reg.Designation
}

// ArithmeticFormat returns the args in an arithmetic format
func (args Args) ArithmeticFormat() ArgsArithmetic {
	return ArgsArithmetic{
		Source1:     args.registerExtract(0xf00000),
		Source2:     args.registerExtract(0x0f0000),
		Destination: args.registerExtract(0x00f000),
	}
}

// ASM returns the representation in assembly language
func (args ArgsArithmetic) ASM() string {
	return fmt.Sprintf("%s %s %s", args.Source1.ASM(), args.Source2.ASM(), args.Destination.ASM())
}

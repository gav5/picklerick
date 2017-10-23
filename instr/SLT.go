package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// SLT sets the destination register to 1 if the first source register is less than the
// branch register; otherwise, it sets the destination register to 0
type SLT struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given SLT instruction
func (i SLT) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i SLT) ASM() string {
	return fmt.Sprintf("SLT %s %s %s", i.args.Destination.ASM(), i.args.Source1.ASM(), i.args.Source2.ASM())
}

// MakeSLT makes an SLT instruction for the given args
func MakeSLT(args instrType.Args) instrType.Base {
	return SLT{args: args.ArithmeticFormat()}
}

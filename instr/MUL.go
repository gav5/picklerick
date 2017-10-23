package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// MUL multiplies the content of two source registers into the destination register
type MUL struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given MUL instruction
func (i MUL) Exec(state *cpuType.State) {
	source1 := state.GetReg(i.args.Source1)
	source2 := state.GetReg(i.args.Source2)
	state.SetReg(i.args.Destination, source1*source2)
}

// ASM returns the representation in assembly language
func (i MUL) ASM() string {
	return fmt.Sprintf("MUL %s", i.args.ASM())
}

// MakeMUL makes a MUL instruction for the given args
func MakeMUL(args instrType.Args) instrType.Base {
	return MUL{args: args.ArithmeticFormat()}
}

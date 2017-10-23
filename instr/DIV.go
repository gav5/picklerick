package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// DIV divides the content of two source registers into the destination register
type DIV struct {
	args instrType.ArgsArithmetic
}

// Exec runs the DIV instruction
func (i DIV) Exec(state *cpuType.State) {
	source1 := state.GetReg(i.args.Source1)
	source2 := state.GetReg(i.args.Source2)
	state.SetReg(i.args.Destination, source1/source2)
}

// ASM returns the representation in assembly language
func (i DIV) ASM() string {
	return fmt.Sprintf("DIV %s", i.args.ASM())
}

// MakeDIV makes a DIV instruction for the given arguments
func MakeDIV(args instrType.Args) instrType.Base {
	return DIV{args: args.ArithmeticFormat()}
}

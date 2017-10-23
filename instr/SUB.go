package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// SUB subtracts the contents of the two source registers into the destination register
type SUB struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given SUB instruction
func (i SUB) Exec(state *cpuType.State) {
	source1 := state.GetReg(i.args.Source1)
	source2 := state.GetReg(i.args.Source2)
	state.SetReg(i.args.Destination, source1-source2)
}

// ASM returns the representation in assembly language
func (i SUB) ASM() string {
	return fmt.Sprintf("SUB %s", i.args.ASM())
}

// MakeSUB makes a SUB instruction for the given args
func MakeSUB(args instrType.Args) instrType.Base {
	return SUB{args: args.ArithmeticFormat()}
}

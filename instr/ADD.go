package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// ADD adds the contents of the two source registers into a destination register
type ADD struct {
	args instrType.ArgsArithmetic
}

// Exec runs the ADD instruction
func (i ADD) Exec(state *cpuType.State) {
	source1 := state.GetReg(i.args.Source1)
	source2 := state.GetReg(i.args.Source2)
	state.SetReg(i.args.Destination, source1+source2)
}

// ASM returns the representation in assembly language
func (i ADD) ASM() string {
	return fmt.Sprintf("ADD %s", i.args.ASM())
}

// MakeADD makes an ADD instruction for the given args
func MakeADD(args instrType.Args) instrType.Base {
	return ADD{args: args.ArithmeticFormat()}
}

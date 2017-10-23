package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// AND logical AND's the contents of two source registers into a desgination register
type AND struct {
	args instrType.ArgsArithmetic
}

// Exec runs the AND instruction
func (i AND) Exec(state *cpuType.State) {
	source1 := state.GetReg(i.args.Source1)
	source2 := state.GetReg(i.args.Source2)
	state.SetReg(i.args.Destination, source1&source2)
}

// ASM returns the representation in assembly language
func (i AND) ASM() string {
	return fmt.Sprintf("AND %s", i.args.ASM())
}

// MakeAND makes an AND instruction for the given args
func MakeAND(args instrType.Args) instrType.Base {
	return AND{args: args.ArithmeticFormat()}
}

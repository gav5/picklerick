package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// OR logical OR's the contents of the two source registers into a destination register
type OR struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given OR instruction
func (i OR) Exec(state *cpuType.State) {
	source1 := state.GetRegBool(i.args.Source1)
	source2 := state.GetRegBool(i.args.Source2)
	state.SetRegBool(i.args.Destination, source1 || source2)
}

// ASM returns the representation in assembly language
func (i OR) ASM() string {
	return fmt.Sprintf("OR %s", i.args.ASM())
}

// MakeOR makes an OR instruction for the given args
func MakeOR(args instrType.Args) instrType.Base {
	return OR{args: args.ArithmeticFormat()}
}

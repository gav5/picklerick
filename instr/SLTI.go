package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// SLTI sets the destination register to 1 if the first source register is less than a
// data value; otherwise, it sets the destination register to 0
type SLTI struct {
	args instrType.ArgsBranch
}

// Exec runs the given SLTI instruction
func (i SLTI) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i SLTI) ASM() string {
	return fmt.Sprintf("SLTI %s", i.args.ASM())
}

// MakeSLTI makes an SLTI instruction for the given args
func MakeSLTI(args instrType.Args) instrType.Base {
	return SLTI{args: args.BranchFormat()}
}

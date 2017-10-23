package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// MOVI transfers address/data directly into a register
type MOVI struct {
	args instrType.ArgsBranch
}

// Exec runs the given MOVI instruction
func (i MOVI) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i MOVI) ASM() string {
	return fmt.Sprintf("MOVI %s", i.args.ASM())
}

// MakeMOVI makes a MOVI instruction for the given args
func MakeMOVI(args instrType.Args) instrType.Base {
	return MOVI{args: args.BranchFormat()}
}

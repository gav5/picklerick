package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// LDI loads some data/address directly to the contents of a register
type LDI struct {
	args instrType.ArgsBranch
}

// Exec runs the given LDI instruction
func (i LDI) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i LDI) ASM() string {
	return fmt.Sprintf("LDI %s", i.args.ASM())
}

// MakeLDI makes an LDI instruction for the given args
func MakeLDI(args instrType.Args) instrType.Base {
	return LDI{args: args.BranchFormat()}
}

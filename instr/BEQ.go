package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BEQ branches to an address when the contents of the branch and destination registers
// are equal to one another
type BEQ struct {
	args instrType.ArgsBranch
}

// Exec runs the BEQ instruction
func (i BEQ) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i BEQ) ASM() string {
	return fmt.Sprintf("BEQ %s", i.args.ASM())
}

// MakeBEQ makes the BEQ instruction for the given args
func MakeBEQ(args instrType.Args) instrType.Base {
	return BEQ{args: args.BranchFormat()}
}

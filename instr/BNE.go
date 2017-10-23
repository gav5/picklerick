package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BNE branches to an address when the content of the branch and destination registers
// are not equal to one another
type BNE struct {
	args instrType.ArgsBranch
}

// Exec runs the BNE instruction
func (i BNE) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i BNE) ASM() string {
	if i.args.Base == 0x0 {
		return fmt.Sprintf("BNE %s %s", i.args.Destination.ASM(), i.args.Address.Dec())
	}
	return fmt.Sprintf("BNE %s %s 0x%04X", i.args.Base.ASM(), i.args.Destination.ASM(), uint32(i.args.Address))
}

// MakeBNE makes a BNE instruction for the given args
func MakeBNE(args instrType.Args) instrType.Base {
	return BNE{args: args.BranchFormat()}
}

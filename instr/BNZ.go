package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BNZ branches to an address when the contents of the branch register is not zero
type BNZ struct {
	args instrType.ArgsBranch
}

// Exec runs the BNZ instruction
func (i BNZ) Exec(state *cpuType.State) {
	base := state.GetReg(i.args.Base)
	if base != 0 {
		addr := uint32(i.args.Address)
		state.SetPC(addr)
	}
}

// ASM returns the representation in assembly language
func (i BNZ) ASM() string {
	return fmt.Sprintf("BNZ %s", i.args.ASM())
}

// MakeBNZ makes a BNZ instruction for the given args
func MakeBNZ(args instrType.Args) instrType.Base {
	return BNZ{args: args.BranchFormat()}
}

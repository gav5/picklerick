package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BLZ branches to an address when the content of the branch register is less than 0
type BLZ struct {
	args instrType.ArgsBranch
}

// Exec runs the BLZ instruction
func (i BLZ) Exec(state *cpuType.State) {
	base := state.GetReg(i.args.Base)
	// NOTE: this should never happen because this is an unsigned value!!
	if base < 0 {
		addr := uint32(i.args.Address)
		state.SetPC(addr)
	}
}

// ASM returns the representation in assembly language
func (i BLZ) ASM() string {
	return fmt.Sprintf("BLZ %s", i.args.ASM())
}

// MakeBLZ makes the BLZ instruction for the given args
func MakeBLZ(args instrType.Args) instrType.Base {
	return BLZ{args: args.BranchFormat()}
}

package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BEZ branches to an address when the contents of the branch register equals zero
type BEZ struct {
	args instrType.ArgsBranch
}

// Exec run the BEZ instruction
func (i BEZ) Exec(state *cpuType.State) {
	base := state.GetReg(i.args.Base)
	if base == 0 {
		addr := uint32(i.args.Address)
		state.SetPC(addr)
	}
}

// ASM returns the representation in assembly language
func (i BEZ) ASM() string {
	return fmt.Sprintf("BEZ %s", i.args.ASM())
}

// MakeBEZ makes the BEZ instruction for the given args
func MakeBEZ(args instrType.Args) instrType.Base {
	return BEZ{args: args.BranchFormat()}
}

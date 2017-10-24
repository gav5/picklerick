package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
	"../reg"
)

// LDI loads some data/address directly to the contents of a register
type LDI struct {
	args instrType.ArgsBranch
}

// Exec runs the given LDI instruction
func (i LDI) Exec(state *cpuType.State) {
	addr := uint32(i.args.Address)
	if addr == 0 {
		base := state.GetReg(i.args.Base)
		state.SetReg(i.args.Destination, base)
	} else {
		state.SetReg(i.args.Destination, reg.Storage(addr))
	}
}

// ASM returns the representation in assembly language
func (i LDI) ASM() string {
	return fmt.Sprintf("LDI %s", i.args.ASM())
}

// MakeLDI makes an LDI instruction for the given args
func MakeLDI(args instrType.Args) instrType.Base {
	return LDI{args: args.BranchFormat()}
}

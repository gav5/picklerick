package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// DIVI divides a data value directly with the content of a register
type DIVI struct {
	args instrType.ArgsBranch
}

// Exec runs the DIVI instruction
func (i DIVI) Exec(state *cpuType.State) {
	base := state.GetReg(i.args.Base)
	dest := state.GetReg(i.args.Destination)
	state.SetReg(i.args.Destination, base/dest)
}

// ASM returns the representation in assembly language
func (i DIVI) ASM() string {
	return fmt.Sprintf("DIVI %s", i.args.ASM())
}

// MakeDIVI makes a DIVI instruction for the given args
func MakeDIVI(args instrType.Args) instrType.Base {
	return DIVI{args: args.BranchFormat()}
}

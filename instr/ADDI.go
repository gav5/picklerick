package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// ADDI adds a data value directly to the contents of a register
type ADDI struct {
	args instrType.ArgsBranch
}

// Exec runs the ADDI instruction
func (i ADDI) Exec(state *cpuType.State) {
	base := state.GetReg(i.args.Base)
	dest := state.GetReg(i.args.Destination)
	state.SetReg(i.args.Destination, base+dest)
}

// ASM returns the representation in assembly language
func (i ADDI) ASM() string {
	return fmt.Sprintf("ADDI %s", i.args.ASM())
}

// MakeADDI makes an ADDI instruction for the given args
func MakeADDI(args instrType.Args) instrType.Base {
	return ADDI{args: args.BranchFormat()}
}

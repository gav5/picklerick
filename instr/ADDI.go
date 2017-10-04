package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// ADDI adds a data value directly to the contents of a register
type ADDI struct {
	args instrType.ArgsBranch
}

// Exec runs the ADDI instruction
func (i ADDI) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i ADDI) ASM() string {
	return fmt.Sprintf("ADDI %s", i.args.ASM())
}

// MakeADDI makes an ADDI instruction for the given args
func MakeADDI(args instrType.Args) instrType.Base {
	return ADDI{args: args.BranchFormat()}
}

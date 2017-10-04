package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// DIVI divides a data value directly with the content of a register
type DIVI struct {
	args instrType.ArgsBranch
}

// Exec runs the DIVI instruction
func (i DIVI) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i DIVI) ASM() string {
	return fmt.Sprintf("DIVI %s", i.args.ASM())
}

// MakeDIVI makes a DIVI instruction for the given args
func MakeDIVI(args instrType.Args) instrType.Base {
	return DIVI{args: args.BranchFormat()}
}

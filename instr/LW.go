package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// LW loads the content of an address into a register
type LW struct {
	args instrType.ArgsBranch
}

// Exec runs the given LW instruction
func (i LW) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i LW) ASM() string {
	return fmt.Sprintf("LW %s", i.args.ASM())
}

// MakeLW makes an LW instruction for the given args
func MakeLW(args instrType.Args) instrType.Base {
	return LW{args: args.BranchFormat()}
}

package main

import "fmt"

// InstructionLDI loads some data/address directly to the contents of a register
type InstructionLDI struct {
	args InstructionArgsBranch
}

func (iLDI InstructionLDI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iLDI InstructionLDI) ASM() string {
	return fmt.Sprintf("LDI %s", iLDI.args.ASM())
}

func makeInstructionLDI(args InstructionArgs) InstructionBase {
	return InstructionLDI{args: args.branchFormat()}
}

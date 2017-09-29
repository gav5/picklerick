package main

import "fmt"

// InstructionADDI adds a data value directly to the contents of a register
type InstructionADDI struct {
	args InstructionArgsBranch
}

func (iADDI InstructionADDI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iADDI InstructionADDI) ASM() string {
	return fmt.Sprintf("ADDI %s", iADDI.args.ASM())
}

func makeInstructionADDI(args InstructionArgs) InstructionBase {
	return InstructionADDI{args: args.branchFormat()}
}

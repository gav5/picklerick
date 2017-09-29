package main

import "fmt"

// InstructionBEZ branches to an address when the contents of the branch register equals zero
type InstructionBEZ struct {
	args InstructionArgsBranch
}

func (iBEZ InstructionBEZ) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBEZ InstructionBEZ) ASM() string {
	return fmt.Sprintf("BEZ %s", iBEZ.args.ASM())
}

func makeInstructionBEZ(args InstructionArgs) InstructionBase {
	return InstructionBEZ{args: args.branchFormat()}
}

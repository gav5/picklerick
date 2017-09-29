package main

import "fmt"

// InstructionBEQ branches to an address when the contents of the branch and destination registers
// are equal to one another
type InstructionBEQ struct {
	args InstructionArgsBranch
}

func (iBEQ InstructionBEQ) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBEQ InstructionBEQ) ASM() string {
	return fmt.Sprintf("BEQ %s", iBEQ.args.ASM())
}

func makeInstructionBEQ(args InstructionArgs) InstructionBase {
	return InstructionBEQ{args: args.branchFormat()}
}

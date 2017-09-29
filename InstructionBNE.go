package main

import "fmt"

// InstructionBNE branches to an address when the content of the branch and destination registers
// are not equal to one another
type InstructionBNE struct {
	args InstructionArgsBranch
}

func (iBNE InstructionBNE) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBNE InstructionBNE) ASM() string {
	return fmt.Sprintf("BNE %s", iBNE.args.ASM())
}

func makeInstructionBNE(args InstructionArgs) InstructionBase {
	return InstructionBNE{args: args.branchFormat()}
}

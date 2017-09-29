package main

import "fmt"

// InstructionBNZ branches to an address when the contents of the branch register is not zero
type InstructionBNZ struct {
	args InstructionArgsBranch
}

func (iBNZ InstructionBNZ) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBNZ InstructionBNZ) ASM() string {
	return fmt.Sprintf("BNZ %s", iBNZ.args.ASM())
}

func makeInstructionBNZ(args InstructionArgs) InstructionBase {
	return InstructionBNZ{args: args.branchFormat()}
}

package main

import "fmt"

// InstructionBGZ branches to an address when the contents of the branch register is greater than 0
type InstructionBGZ struct {
	args InstructionArgsBranch
}

func (iBGZ InstructionBGZ) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBGZ InstructionBGZ) ASM() string {
	return fmt.Sprintf("BGZ %s", iBGZ.args.ASM())
}

func makeInstructionBGZ(args InstructionArgs) InstructionBase {
	return InstructionBGZ{args: args.branchFormat()}
}

package main

import "fmt"

// InstructionMULI multiplies a data value directly with the content of a register
type InstructionMULI struct {
	args InstructionArgsBranch
}

func (iMULI InstructionMULI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iMULI InstructionMULI) ASM() string {
	return fmt.Sprintf("MULI %s", iMULI.args.ASM())
}

func makeInstructionMULI(args InstructionArgs) InstructionBase {
	return InstructionMULI{args: args.branchFormat()}
}

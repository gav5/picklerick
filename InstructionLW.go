package main

import "fmt"

// InstructionLW loads the content of an address into a register
type InstructionLW struct {
	args InstructionArgsBranch
}

func (iLW InstructionLW) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iLW InstructionLW) ASM() string {
	return fmt.Sprintf("LW %s", iLW.args.ASM())
}

func makeInstructionLW(args InstructionArgs) InstructionBase {
	return InstructionLW{args: args.branchFormat()}
}

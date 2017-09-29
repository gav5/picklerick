package main

import "fmt"

// InstructionDIVI divides a data value directly with the content of a register
type InstructionDIVI struct {
	args InstructionArgsBranch
}

func (iDIVI InstructionDIVI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iDIVI InstructionDIVI) ASM() string {
	return fmt.Sprintf("DIVI %s", iDIVI.args.ASM())
}

func makeInstructionDIVI(args InstructionArgs) InstructionBase {
	return InstructionDIVI{args: args.branchFormat()}
}

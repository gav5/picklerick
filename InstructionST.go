package main

import "fmt"

// InstructionST stores the content of a register into an address
type InstructionST struct {
	args InstructionArgsBranch
}

func (iST InstructionST) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iST InstructionST) ASM() string {
	return fmt.Sprintf("ST %s", iST.args.ASM())
}

func makeInstructionST(args InstructionArgs) InstructionBase {
	return InstructionST{args: args.branchFormat()}
}

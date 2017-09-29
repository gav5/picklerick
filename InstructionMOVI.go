package main

import "fmt"

// InstructionMOVI transfers address/data directly into a register
type InstructionMOVI struct {
	args InstructionArgsBranch
}

func (iMOVI InstructionMOVI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iMOVI InstructionMOVI) ASM() string {
	return fmt.Sprintf("MOVI %s", iMOVI.args.ASM())
}

func makeInstructionMOVI(args InstructionArgs) InstructionBase {
	return InstructionMOVI{args: args.branchFormat()}
}

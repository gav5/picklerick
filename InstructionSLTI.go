package main

import "fmt"

// InstructionSLTI sets the destination register to 1 if the first source register is less than a
// data value; otherwise, it sets the destination register to 0
type InstructionSLTI struct {
	args InstructionArgsBranch
}

func (iSTLI InstructionSLTI) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iSTLI InstructionSLTI) ASM() string {
	return fmt.Sprintf("SLTI %s", iSTLI.args.ASM())
}

func makeInstructionSLTI(args InstructionArgs) InstructionBase {
	return InstructionSLTI{args: args.branchFormat()}
}

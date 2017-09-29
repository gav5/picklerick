package main

import "fmt"

// InstructionBLZ branches to an address when the content of the branch register is less than 0
type InstructionBLZ struct {
	args InstructionArgsBranch
}

func (iBLZ InstructionBLZ) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iBLZ InstructionBLZ) ASM() string {
	return fmt.Sprintf("BLZ %s", iBLZ.args.ASM())
}

func makeInstructionBLZ(args InstructionArgs) InstructionBase {
	return InstructionBLZ{args: args.branchFormat()}
}

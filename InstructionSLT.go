package main

import "fmt"

// InstructionSLT sets the destination register to 1 if the first source register is less than the
// branch register; otherwise, it sets the destination register to 0
type InstructionSLT struct {
	args InstructionArgsArithmetic
}

func (iSLT InstructionSLT) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iSLT InstructionSLT) ASM() string {
	return fmt.Sprintf("SLT %s", iSLT.args.ASM())
}

func makeInstructionSLT(args InstructionArgs) InstructionBase {
	return InstructionSLT{args: args.arithmeticFormat()}
}

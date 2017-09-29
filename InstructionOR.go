package main

import "fmt"

// InstructionOR logical OR's the contents of the two source registers into a destination register
type InstructionOR struct {
	args InstructionArgsArithmetic
}

func (iOR InstructionOR) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iOR InstructionOR) ASM() string {
	return fmt.Sprintf("OR %s", iOR.args.ASM())
}

func makeInstructionOR(args InstructionArgs) InstructionBase {
	return InstructionOR{args: args.arithmeticFormat()}
}

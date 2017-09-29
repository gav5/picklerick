package main

import "fmt"

// InstructionMUL multiplies the content of two source registers into the destination register
type InstructionMUL struct {
	args InstructionArgsArithmetic
}

func (iMUL InstructionMUL) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iMUL InstructionMUL) ASM() string {
	return fmt.Sprintf("MUL %s", iMUL.args.ASM())
}

func makeInstructionMUL(args InstructionArgs) InstructionBase {
	return InstructionMUL{args: args.arithmeticFormat()}
}

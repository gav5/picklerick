package main

import "fmt"

// InstructionADD adds the contents of the two source registers into a destination register
type InstructionADD struct {
	args InstructionArgsArithmetic
}

func (iADD InstructionADD) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iADD InstructionADD) ASM() string {
	return fmt.Sprintf("ADD %s", iADD.args.ASM())
}

func makeInstructionADD(args InstructionArgs) InstructionBase {
	return InstructionADD{args: args.arithmeticFormat()}
}

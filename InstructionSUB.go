package main

import "fmt"

// InstructionSUB subtracts the contents of the two source registers into the destination register
type InstructionSUB struct {
	args InstructionArgsArithmetic
}

func (iSUB InstructionSUB) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iSUB InstructionSUB) ASM() string {
	return fmt.Sprintf("SUB %s", iSUB.args.ASM())
}

func makeInstructionSUB(args InstructionArgs) InstructionBase {
	return InstructionSUB{args: args.arithmeticFormat()}
}

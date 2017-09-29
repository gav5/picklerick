package main

import "fmt"

// InstructionAND logical AND's the contents of two source registers into a desgination register
type InstructionAND struct {
	args InstructionArgsArithmetic
}

func (iAND InstructionAND) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iAND InstructionAND) ASM() string {
	return fmt.Sprintf("AND %s", iAND.args.ASM())
}

func makeInstructionAND(args InstructionArgs) InstructionBase {
	return InstructionAND{args: args.arithmeticFormat()}
}

package main

import "fmt"

// InstructionDIV divides the content of two source registers into the destination register
type InstructionDIV struct {
	args InstructionArgsArithmetic
}

func (iDIV InstructionDIV) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iDIV InstructionDIV) ASM() string {
	return fmt.Sprintf("DIV %s", iDIV.args.ASM())
}

func makeInstructionDIV(args InstructionArgs) InstructionBase {
	return InstructionDIV{args: args.arithmeticFormat()}
}

package main

import "fmt"

// InstructionMOV transfers the contents of one register into another
type InstructionMOV struct {
	args InstructionArgsArithmetic
}

func (iMOV InstructionMOV) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iMOV InstructionMOV) ASM() string {
	return fmt.Sprintf("MOV %s", iMOV.args.ASM())
}

func makeInstructionMOV(args InstructionArgs) InstructionBase {
	return InstructionMOV{args: args.arithmeticFormat()}
}

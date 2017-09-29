package main

import "fmt"

// InstructionRD reads the content of the I/P buffer into the accumulator
type InstructionRD struct {
	args InstructionArgsIO
}

func (iRD InstructionRD) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iRD InstructionRD) ASM() string {
	return fmt.Sprintf("RD %s", iRD.args.ASM())
}

func makeInstructionRD(args InstructionArgs) InstructionBase {
	return InstructionRD{args: args.ioFormat()}
}

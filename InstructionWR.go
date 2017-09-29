package main

import "fmt"

// InstructionWR writes the content of the accumulator into the O/P buffer
type InstructionWR struct {
	args InstructionArgsIO
}

func (iWR InstructionWR) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iWR InstructionWR) ASM() string {
	return fmt.Sprintf("WR %s", iWR.args.ASM())
}

func makeInstructionWR(args InstructionArgs) InstructionBase {
	return InstructionWR{args: args.ioFormat()}
}

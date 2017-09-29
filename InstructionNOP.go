package main

// InstructionNOP does nothing and moves to the next instruction
type InstructionNOP struct {
	// no arguments needed
}

func (iNOP InstructionNOP) exec(pcb PCB) PCB {
	// this does nothing, so return the same state
	return pcb
}

// ASM returns the representation in assembly language
func (iNOP InstructionNOP) ASM() string {
	return "NOP"
}

func makeInstructionNOP(args InstructionArgs) InstructionBase {
	return InstructionNOP{}
}

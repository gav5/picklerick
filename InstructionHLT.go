package main

// InstructionHLT signals the logical end of a program (i.e. it "halts" the program)
type InstructionHLT struct {
}

func (iHLT InstructionHLT) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iHLT InstructionHLT) ASM() string {
	return "HLT"
}

func makeInstructionHLT(args InstructionArgs) InstructionBase {
	return InstructionHLT{}
}

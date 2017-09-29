package main

// InstructionBase describes the base requirements for a CPU instruction
type InstructionBase interface {
	exec(pcb PCB) PCB
	// ASM returns the representation in assembly language
	ASM() string
}

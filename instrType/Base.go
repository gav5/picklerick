package instrType

import "../proc"

// Base describes the base requirements for a CPU instruction
type Base interface {
	Exec(pcb proc.PCB) proc.PCB
	// ASM returns the representation in assembly language
	ASM() string
}

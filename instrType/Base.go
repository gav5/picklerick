package instrType

import "../cpuType"

// Base describes the base requirements for a CPU instruction
type Base interface {
	Exec(state *cpuType.State)
	// ASM returns the representation in assembly language
	ASM() string
}

package ivm

// Instruction is the interface for an instruction to the CPU core.
type Instruction interface {
	Execute(ip InstructionProxy)
	Assembly() string
}

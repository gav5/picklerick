package instr

import (
	"fmt"

	"../ivm"
)

// JMP jumps to a specified location
type JMP struct {
	args ivm.InstructionArgsJump
}

// Execute runs the JMP instruction
func (i JMP) Execute(ip ivm.InstructionProxy) {
	ip.SetProgramCounter(i.args.Address)
}

// Assembly returns the representation in assembly language
func (i JMP) Assembly() string {
	return fmt.Sprintf("JMP %s", i.args.ASM())
}

// MakeJMP makes a JMP instruction for the given args
func MakeJMP(args ivm.InstructionArgs) ivm.Instruction {
	return JMP{args: args.JumpFormat()}
}

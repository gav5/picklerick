package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// JMP jumps to a specified location
type JMP struct {
	args instrType.ArgsJump
}

// Exec runs the JMP instruction
func (i JMP) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i JMP) ASM() string {
	return fmt.Sprintf("JMP %s", i.args.ASM())
}

// MakeJMP makes a JMP instruction for the given args
func MakeJMP(args instrType.Args) instrType.Base {
	return JMP{args: args.JumpFormat()}
}

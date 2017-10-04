package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// MUL multiplies the content of two source registers into the destination register
type MUL struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given MUL instruction
func (i MUL) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i MUL) ASM() string {
	return fmt.Sprintf("MUL %s", i.args.ASM())
}

// MakeMUL makes a MUL instruction for the given args
func MakeMUL(args instrType.Args) instrType.Base {
	return MUL{args: args.ArithmeticFormat()}
}

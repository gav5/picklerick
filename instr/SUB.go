package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// SUB subtracts the contents of the two source registers into the destination register
type SUB struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given SUB instruction
func (i SUB) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i SUB) ASM() string {
	return fmt.Sprintf("SUB %s", i.args.ASM())
}

// MakeSUB makes a SUB instruction for the given args
func MakeSUB(args instrType.Args) instrType.Base {
	return SUB{args: args.ArithmeticFormat()}
}

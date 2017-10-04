package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// DIV divides the content of two source registers into the destination register
type DIV struct {
	args instrType.ArgsArithmetic
}

// Exec runs the DIV instruction
func (i DIV) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i DIV) ASM() string {
	return fmt.Sprintf("DIV %s", i.args.ASM())
}

// MakeDIV makes a DIV instruction for the given arguments
func MakeDIV(args instrType.Args) instrType.Base {
	return DIV{args: args.ArithmeticFormat()}
}

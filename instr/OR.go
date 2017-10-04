package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// OR logical OR's the contents of the two source registers into a destination register
type OR struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given OR instruction
func (i OR) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i OR) ASM() string {
	return fmt.Sprintf("OR %s", i.args.ASM())
}

// MakeOR makes an OR instruction for the given args
func MakeOR(args instrType.Args) instrType.Base {
	return OR{args: args.ArithmeticFormat()}
}

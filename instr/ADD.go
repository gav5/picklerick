package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// ADD adds the contents of the two source registers into a destination register
type ADD struct {
	args instrType.ArgsArithmetic
}

// Exec runs the ADD instruction
func (i ADD) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i ADD) ASM() string {
	return fmt.Sprintf("ADD %s", i.args.ASM())
}

// MakeADD makes an ADD instruction for the given args
func MakeADD(args instrType.Args) instrType.Base {
	return ADD{args: args.ArithmeticFormat()}
}

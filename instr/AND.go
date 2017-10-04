package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// AND logical AND's the contents of two source registers into a desgination register
type AND struct {
	args instrType.ArgsArithmetic
}

// Exec runs the AND instruction
func (i AND) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i AND) ASM() string {
	return fmt.Sprintf("AND %s", i.args.ASM())
}

// MakeAND makes an AND instruction for the given args
func MakeAND(args instrType.Args) instrType.Base {
	return AND{args: args.ArithmeticFormat()}
}

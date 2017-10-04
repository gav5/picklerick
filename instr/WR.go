package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// WR writes the content of the accumulator into the O/P buffer
type WR struct {
	args instrType.ArgsIO
}

// Exec runs the given WR instruction
func (i WR) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i WR) ASM() string {
	return fmt.Sprintf("WR %s", i.args.ASM())
}

// MakeWR makes a WR instruction for the given args
func MakeWR(args instrType.Args) instrType.Base {
	return WR{args: args.IOFormat()}
}

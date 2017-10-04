package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// RD reads the content of the I/P buffer into the accumulator
type RD struct {
	args instrType.ArgsIO
}

// Exec runs the given RD instruction
func (i RD) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i RD) ASM() string {
	return fmt.Sprintf("RD %s", i.args.ASM())
}

// MakeRD makes an RD instruction for the given args
func MakeRD(args instrType.Args) instrType.Base {
	return RD{args: args.IOFormat()}
}

package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
	"../reg"
)

// RD reads the content of the I/P buffer into the accumulator
type RD struct {
	args instrType.ArgsIO
}

// Exec runs the given RD instruction
func (i RD) Exec(state *cpuType.State) {
	addr := uint32(i.args.Address)
	val := reg.Storage(state.GetAddr(addr))
	state.SetReg(i.args.Register1, val)
}

// ASM returns the representation in assembly language
func (i RD) ASM() string {
	return fmt.Sprintf("RD %s", i.args.ASM())
}

// MakeRD makes an RD instruction for the given args
func MakeRD(args instrType.Args) instrType.Base {
	return RD{args: args.IOFormat()}
}

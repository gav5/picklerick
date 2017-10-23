package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// WR writes the content of the accumulator into the O/P buffer
type WR struct {
	args instrType.ArgsIO
}

// Exec runs the given WR instruction
func (i WR) Exec(state *cpuType.State) {
	addr := uint32(i.args.Address)
	r1 := state.GetReg(i.args.Register1)
	state.SetAddr(addr, uint32(r1))
}

// ASM returns the representation in assembly language
func (i WR) ASM() string {
	return fmt.Sprintf("WR %s", i.args.ASM())
}

// MakeWR makes a WR instruction for the given args
func MakeWR(args instrType.Args) instrType.Base {
	return WR{args: args.IOFormat()}
}

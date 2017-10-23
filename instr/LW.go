package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// LW loads the content of an address into a register
type LW struct {
	args instrType.ArgsBranch
}

// Exec runs the given LW instruction
func (i LW) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i LW) ASM() string {
	return fmt.Sprintf("LW %s %s(%s)", i.args.Destination.ASM(), i.args.Address.Dec(), i.args.Base.ASM())
}

// MakeLW makes an LW instruction for the given args
func MakeLW(args instrType.Args) instrType.Base {
	return LW{args: args.BranchFormat()}
}

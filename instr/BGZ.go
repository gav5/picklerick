package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// BGZ branches to an address when the contents of the branch register is greater than 0
type BGZ struct {
	args instrType.ArgsBranch
}

// Exec runs the BGZ instruction
func (i BGZ) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i BGZ) ASM() string {
	return fmt.Sprintf("BGZ %s", i.args.ASM())
}

// MakeBGZ make the BGZ instruction for the given args
func MakeBGZ(args instrType.Args) instrType.Base {
	return BGZ{args: args.BranchFormat()}
}

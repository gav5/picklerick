package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// JMP jumps to a specified location
type JMP struct {
	args instrType.ArgsJump
}

// Exec runs the JMP instruction
func (i JMP) Exec(state *cpuType.State) {
	pc := uint32(i.args.Address)
	state.SetPC(pc)
}

// ASM returns the representation in assembly language
func (i JMP) ASM() string {
	return fmt.Sprintf("JMP %s", i.args.ASM())
}

// MakeJMP makes a JMP instruction for the given args
func MakeJMP(args instrType.Args) instrType.Base {
	return JMP{args: args.JumpFormat()}
}

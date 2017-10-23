package instr

import (
	"fmt"

	"../cpuType"
	"../instrType"
)

// MOV transfers the contents of one register into another
type MOV struct {
	args instrType.ArgsArithmetic
}

// Exec runs the given MOV instruction
func (i MOV) Exec(state *cpuType.State) {
	// TODO: make this actually do what it's supposed to do
}

// ASM returns the representation in assembly language
func (i MOV) ASM() string {
	return fmt.Sprintf("MOV %s", i.args.ASM())
}

// MakeMOV makes an MOV instruction for the given args
func MakeMOV(args instrType.Args) instrType.Base {
	return MOV{args: args.ArithmeticFormat()}
}

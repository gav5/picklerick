package instr

import (
	"fmt"

	"../ivm"
)

// MOV transfers the contents of one register into another
type MOV struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the given MOV instruction
func (i MOV) Execute(ip ivm.InstructionProxy) {
	source := ip.RegisterWord(i.args.Source2)
	ip.SetRegisterWord(i.args.Source1, source)
}

// Assembly returns the representation in assembly language
func (i MOV) Assembly() string {
	// return fmt.Sprintf("MOV %s", i.args.ASM())
	return fmt.Sprintf("MOV %s %s", i.args.Source1, i.args.Source2)
}

// MakeMOV makes an MOV instruction for the given args
func MakeMOV(args ivm.InstructionArgs) ivm.Instruction {
	return MOV{args: args.ArithmeticFormat()}
}

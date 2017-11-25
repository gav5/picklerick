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
	source := ip.RegisterUint32(i.args.Source1)
	ip.SetRegisterUint32(i.args.Destination, source)
}

// Assembly returns the representation in assembly language
func (i MOV) Assembly() string {
	return fmt.Sprintf("MOV %s", i.args.ASM())
}

// MakeMOV makes an MOV instruction for the given args
func MakeMOV(args ivm.InstructionArgs) ivm.Instruction {
	return MOV{args: args.ArithmeticFormat()}
}

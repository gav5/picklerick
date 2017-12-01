package instr

import (
	"fmt"

	"../ivm"
)

// RD reads the content of the I/P buffer into the accumulator
type RD struct {
	args ivm.InstructionArgsIO
}

// Execute runs the given RD instruction
func (i RD) Execute(ip ivm.InstructionProxy) {
	var addr ivm.Address
	if i.args.Register2 == ivm.R0 {
		addr = i.args.Address
	} else {
		addr = ivm.Address(ip.RegisterWord(i.args.Register2))
	}
	val := ip.AddressFetchWord(addr)
	ip.SetRegisterWord(i.args.Register1, val)
}

// Assembly returns the representation in assembly language
func (i RD) Assembly() string {
	return fmt.Sprintf("RD %s", i.args.ASM())
}

// MakeRD makes an RD instruction for the given args
func MakeRD(args ivm.InstructionArgs) ivm.Instruction {
	return RD{args: args.IOFormat()}
}

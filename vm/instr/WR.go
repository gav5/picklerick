package instr

import (
	"fmt"

	"../ivm"
)

// WR writes the content of the accumulator into the O/P buffer
type WR struct {
	args ivm.InstructionArgsIO
}

// Execute runs the given WR instruction
func (i WR) Execute(ip ivm.InstructionProxy) {
	var addr ivm.Address
	if i.args.Register2 == ivm.R0 {
		addr = i.args.Address
	} else {
		addr = ivm.Address(ip.RegisterWord(i.args.Register2))
	}
	r1 := ip.RegisterWord(i.args.Register1)
	ip.AddressWriteWord(addr, r1)
}

// Assembly returns the representation in assembly language
func (i WR) Assembly() string {
	return fmt.Sprintf("WR %s", i.args.ASMHex())
}

// MakeWR makes a WR instruction for the given args
func MakeWR(args ivm.InstructionArgs) ivm.Instruction {
	return WR{args: args.IOFormat()}
}

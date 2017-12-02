package program

import (
	"fmt"

	"../../vm/ivm"
)

// DecodedInstruction holds the human-readable form of an instruction.
type DecodedInstruction struct {
	ADDR ivm.Address
	RAW  uint32
	ASM  string
}

func (di DecodedInstruction) String() string {
	return fmt.Sprintf("%04X | %08X | %s", int(di.ADDR), di.RAW, di.ASM)
}

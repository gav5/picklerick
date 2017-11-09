package vm

import (
	"fmt"

	"./ivm"
)

// UnrecognizedOpcodeError indicates the opcode was not valid
type UnrecognizedOpcodeError struct {
	op ivm.Opcode
}

func (err UnrecognizedOpcodeError) Error() string {
	return fmt.Sprintf("The opcode %v is unrecognized", err.op)
}

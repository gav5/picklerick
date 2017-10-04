package instrDecode

import (
	"fmt"

	"../instr"
	"../instrType"
)

// UnrecognizedOpcodeError indicates the opcode was not valid
type UnrecognizedOpcodeError struct {
	op instrType.Opcode
}

func (err UnrecognizedOpcodeError) Error() string {
	return fmt.Sprintf("The opcode %v is unrecognized", err.op)
}

func decodeOpcode(op instrType.Opcode) (instrType.Factory, error) {
	switch op {
	case 0x00:
		return instr.MakeRD, nil
	case 0x01:
		return instr.MakeWR, nil
	case 0x02:
		return instr.MakeST, nil
	case 0x03:
		return instr.MakeLW, nil
	case 0x04:
		return instr.MakeMOV, nil
	case 0x05:
		return instr.MakeADD, nil
	case 0x06:
		return instr.MakeSUB, nil
	case 0x07:
		return instr.MakeMUL, nil
	case 0x08:
		return instr.MakeDIV, nil
	case 0x09:
		return instr.MakeAND, nil
	case 0x0a:
		return instr.MakeOR, nil
	case 0x0b:
		return instr.MakeMOVI, nil
	case 0x0c:
		return instr.MakeADDI, nil
	case 0x0d:
		return instr.MakeMULI, nil
	case 0x0e:
		return instr.MakeDIVI, nil
	case 0x0f:
		return instr.MakeLDI, nil
	case 0x10:
		return instr.MakeSLT, nil
	case 0x11:
		return instr.MakeSLTI, nil
	case 0x12:
		return instr.MakeHLT, nil
	case 0x13:
		return instr.MakeNOP, nil
	case 0x14:
		return instr.MakeJMP, nil
	case 0x15:
		return instr.MakeBEQ, nil
	case 0x16:
		return instr.MakeBNE, nil
	case 0x17:
		return instr.MakeBEZ, nil
	case 0x18:
		return instr.MakeBNZ, nil
	case 0x19:
		return instr.MakeBGZ, nil
	case 0x1a:
		return instr.MakeBLZ, nil
	default:
		return nil, UnrecognizedOpcodeError{op: op}
	}
}

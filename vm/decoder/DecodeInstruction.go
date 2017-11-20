package decoder

import (
	"../../util"
	"../ivm"
)

// DecodeInstruction makes an instruction from the given 32-bit value
func DecodeInstruction(val uint32) (ivm.Instruction, error) {
	op := ivm.Opcode(util.BitExtract32(val, 0x3f000000))
	args := ivm.InstructionArgs(util.BitExtract32(val, 0x00ffffff))
	factory, err := decodeOpcode(op)
	if err != nil {
		return nil, err
	}
	return factory(args), nil
}

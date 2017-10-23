package decoder

import (
	"../instrType"
	"../util"
)

// InstrFromUint32 makes an instruction from the given 32-bit instruction
func InstrFromUint32(val uint32) (instrType.Base, error) {
	comp := instrType.Components{
		Opcode: instrType.Opcode(util.BitExtract32(val, 0x3f000000)),
		Args:   instrType.Args(util.BitExtract32(val, 0x00ffffff)),
	}
	factory, err := decodeOpcode(comp.Opcode)
	if err != nil {
		return nil, err
	}
	return factory(comp.Args), nil
}

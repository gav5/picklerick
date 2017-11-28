package ivm

// Opcode encapsulates the opcode of a binary instruction
type Opcode uint8

// func makeOpcode(s string) (Opcode, error) {
// 	i, err := util.HexExtract8(s)
// 	if err != nil {
// 		return Opcode(0), err
// 	}
// 	return Opcode(util.BitExtract8(i, 0x3f)), nil
// }

// InstructionFactory describes a function that builds an instruction
// for the given binary components
type InstructionFactory func(InstructionArgs) Instruction

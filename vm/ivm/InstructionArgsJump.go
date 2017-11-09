package ivm

// InstructionArgsJump encapsulates the args for jump instructions
type InstructionArgsJump struct {
	Address Address
}

// JumpFormat returns the args in a jump format
func (args InstructionArgs) JumpFormat() InstructionArgsJump {
	return InstructionArgsJump{
		Address: args.extractAddress(0xffffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsJump) ASM() string {
	return args.Address.Hex()
}

package instrType

import "../bus"

// ArgsJump encapsulates the args for jump instructions
type ArgsJump struct {
	Address bus.Address
}

// JumpFormat returns the args in a jump format
func (args Args) JumpFormat() ArgsJump {
	return ArgsJump{
		Address: args.addressExtract(0xffffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsJump) ASM() string {
	return args.Address.Hex()
}

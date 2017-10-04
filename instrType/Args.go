package instrType

import (
	"fmt"

	"../bus"
	"../reg"
	"../util"
)

// Args encapsulates the args given to a binary instruction
type Args uint32

func makeArgs(s string) (Args, error) {
	i, err := util.HexExtract32(s)
	if err != nil {
		return Args(0), err
	}
	return Args(i), nil
}

// RegisterExtract extracts out a RegisterDesignation for the given mask
func (args Args) RegisterExtract(mask uint32) reg.Designation {
	return reg.Designation(util.BitExtract32(uint32(args), mask))
}

// AddressExtract extracts out a BusAddress for a given mask
func (args Args) AddressExtract(mask uint32) bus.Address {
	return bus.Address(util.BitExtract32(uint32(args), mask))
}

// ArgsArithmetic encapsulates the args for arithmetic instructions
type ArgsArithmetic struct {
	source1     reg.Designation
	source2     reg.Designation
	destination reg.Designation
}

// ArithmeticFormat returns the args in an arithmetic format
func (args Args) ArithmeticFormat() ArgsArithmetic {
	return ArgsArithmetic{
		source1:     args.RegisterExtract(0xf00000),
		source2:     args.RegisterExtract(0x0f0000),
		destination: args.RegisterExtract(0x00f000),
	}
}

// ASM returns the representation in assembly language
func (args ArgsArithmetic) ASM() string {
	return fmt.Sprintf("%s %s %s", args.source1.ASM(), args.source2.ASM(), args.destination.ASM())
}

// ArgsBranch encapsulates the args for branch instructions
type ArgsBranch struct {
	base        reg.Designation
	destination reg.Designation
	address     bus.Address
}

// BranchFormat returns the args in a branch format
func (args Args) BranchFormat() ArgsBranch {
	return ArgsBranch{
		base:        args.RegisterExtract(0xf00000),
		destination: args.RegisterExtract(0x0f0000),
		address:     args.AddressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsBranch) ASM() string {
	if args.base == 0x0 {
		return fmt.Sprintf("%s %s", args.destination.ASM(), args.address.Dec())
	}
	return fmt.Sprintf("%s (%s)", args.destination.ASM(), args.base.ASM())
}

// ArgsJump encapsulates the args for jump instructions
type ArgsJump struct {
	address bus.Address
}

// JumpFormat returns the args in a jump format
func (args Args) JumpFormat() ArgsJump {
	return ArgsJump{
		address: args.AddressExtract(0xffffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsJump) ASM() string {
	return args.address.Hex()
}

// ArgsIO encapsulates the args for IO instructions
type ArgsIO struct {
	register1 reg.Designation
	register2 reg.Designation
	address   bus.Address
}

// IOFormat returns the args in an IO format
func (args Args) IOFormat() ArgsIO {
	return ArgsIO{
		register1: args.RegisterExtract(0xf00000),
		register2: args.RegisterExtract(0x0f0000),
		address:   args.AddressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsIO) ASM() string {
	if args.register2 == 0x0 {
		return fmt.Sprintf("%s %s", args.register1.ASM(), args.address.Hex())
	}
	return fmt.Sprintf("%s (%s)", args.register1.ASM(), args.register2.ASM())
}

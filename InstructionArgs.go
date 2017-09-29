package main

import "fmt"

// InstructionArgs encapsulates the args given to a binary instruction
type InstructionArgs uint32

func makeInstructionArgs(s string) (InstructionArgs, error) {
	i, err := HexExtract32(s)
	if err != nil {
		return InstructionArgs(0), err
	}
	return InstructionArgs(i), nil
}

// RegisterExtract extracts out a RegisterDesignation for the given mask
func (args InstructionArgs) RegisterExtract(mask uint32) RegisterDesignation {
	return RegisterDesignation(BitExtract32(uint32(args), mask))
}

// AddressExtract extracts out a BusAddress for a given mask
func (args InstructionArgs) AddressExtract(mask uint32) BusAddress {
	return BusAddress(BitExtract32(uint32(args), mask))
}

// InstructionArgsArithmetic encapsulates the args for arithmetic instructions
type InstructionArgsArithmetic struct {
	source1     RegisterDesignation
	source2     RegisterDesignation
	destination RegisterDesignation
}

func (args InstructionArgs) arithmeticFormat() InstructionArgsArithmetic {
	return InstructionArgsArithmetic{
		source1:     args.RegisterExtract(0xf00000),
		source2:     args.RegisterExtract(0x0f0000),
		destination: args.RegisterExtract(0x00f000),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsArithmetic) ASM() string {
	return fmt.Sprintf("%s %s %s", args.source1.ASM(), args.source2.ASM(), args.destination.ASM())
}

// InstructionArgsBranch encapsulates the args for branch instructions
type InstructionArgsBranch struct {
	base        RegisterDesignation
	destination RegisterDesignation
	address     BusAddress
}

func (args InstructionArgs) branchFormat() InstructionArgsBranch {
	return InstructionArgsBranch{
		base:        args.RegisterExtract(0xf00000),
		destination: args.RegisterExtract(0x0f0000),
		address:     args.AddressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsBranch) ASM() string {
	if args.base == 0x0 {
		return fmt.Sprintf("%s %s", args.destination.ASM(), args.address.Dec())
	}
	return fmt.Sprintf("%s (%s)", args.destination.ASM(), args.base.ASM())
}

// InstructionArgsJump encapsulates the args for jump instructions
type InstructionArgsJump struct {
	address BusAddress
}

func (args InstructionArgs) jumpFormat() InstructionArgsJump {
	return InstructionArgsJump{
		address: args.AddressExtract(0xffffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsJump) ASM() string {
	return args.address.Hex()
}

// InstructionArgsIO encapsulates the args for IO instructions
type InstructionArgsIO struct {
	register1 RegisterDesignation
	register2 RegisterDesignation
	address   BusAddress
}

func (args InstructionArgs) ioFormat() InstructionArgsIO {
	return InstructionArgsIO{
		register1: args.RegisterExtract(0xf00000),
		register2: args.RegisterExtract(0x0f0000),
		address:   args.AddressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args InstructionArgsIO) ASM() string {
	if args.register2 == 0x0 {
		return fmt.Sprintf("%s %s", args.register1.ASM(), args.address.Hex())
	}
	return fmt.Sprintf("%s (%s)", args.register1.ASM(), args.register2.ASM())
}

package instrType

import (
	"fmt"

	"../bus"
	"../reg"
)

// ArgsBranch encapsulates the args for branch instructions
type ArgsBranch struct {
	Base        reg.Designation
	Destination reg.Designation
	Address     bus.Address
}

// BranchFormat returns the args in a branch format
func (args Args) BranchFormat() ArgsBranch {
	return ArgsBranch{
		Base:        args.registerExtract(0xf00000),
		Destination: args.registerExtract(0x0f0000),
		Address:     args.addressExtract(0x00ffff),
	}
}

// ASM returns the representation in assembly language
func (args ArgsBranch) ASM() string {
	if args.Base == 0x0 {
		return fmt.Sprintf("%s %s", args.Destination.ASM(), args.Address.Dec())
	}
	return fmt.Sprintf("%s (%s)", args.Destination.ASM(), args.Base.ASM())
}

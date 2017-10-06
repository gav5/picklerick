package instrType

import (
	"../bus"
	"../reg"
	"../util"
)

// Args encapsulates the args given to a binary instruction
type Args uint32

func (args Args) registerExtract(mask uint32) reg.Designation {
	return reg.Designation(util.BitExtract32(uint32(args), mask))
}

func (args Args) addressExtract(mask uint32) bus.Address {
	return bus.Address(util.BitExtract32(uint32(args), mask))
}

package ivm

const (
	// RAMNumWords is the number of words in the virtual machine RAM.
	RAMNumWords = 1024

	// RAMNumFrames is the number of frames in the virtual machine RAM.
	RAMNumFrames = RAMNumWords / FrameSize
)

// IRAM is the kernel interface into a virtual machine random access memory.
// (this is to avoid circular dependencies and regulate kernel access)
type IRAM interface {

	// Addresses as gateways to words
	AddressFetchWord(Address) Word
	AddressWriteWord(Address, Word)

	// Addresses as gateways to uint32's
	AddressFetchUint32(Address) uint32
	AddressWriteUint32(Address, uint32)

	// Addresses as gateways to int32's
	AddressFetchInt32(Address) int32
	AddressWriteInt32(Address, int32)

	// Addresses as gateways to bool's
	AddressFetchBool(Address) bool
	AddressWriteBool(Address, bool)

	// Frame management
	FrameFetch(FrameNumber) Frame
	FrameWrite(FrameNumber, Frame)
}

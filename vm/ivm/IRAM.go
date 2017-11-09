package ivm

// IRAM is the kernel interface into a virtual machine random access memory.
// (this is because vm uses the kernel, so we have to avoid circular dependencies)
// (this also helps regulate the access of the vm by the kernel to avoid abuse)
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
}

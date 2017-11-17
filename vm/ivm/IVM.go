package ivm

// NumCores is the number of cores in the virtual machine.
const NumCores = 4

// IVM is the kernel interface into the virtual machine.
// (this is because vm uses the kernel, so we have to avoid circular dependencies)
// (this also helps regulate the access of the vm by the kernel to avoid abuse)
type IVM interface {

	// Core Interfacing (must specify coreNum as well)
	ProgramCounter(uint8) Address
	SetProgramCounter(uint8, Address)
	Halt(uint8)
	ResetCore(uint8)
	RegisterWord(uint8, RegisterDesignation) Word
	SetRegisterWord(uint8, RegisterDesignation, Word)
	RegisterUint32(uint8, RegisterDesignation) uint32
	SetRegisterUint32(uint8, RegisterDesignation, uint32)
	RegisterInt32(uint8, RegisterDesignation) int32
	SetRegisterInt32(uint8, RegisterDesignation, int32)
	RegisterBool(uint8, RegisterDesignation) bool
	SetRegisterBool(uint8, RegisterDesignation, bool)

	// Working with RAM
	RAMAddressFetchWord(Address) Word
	RAMAddressWriteWord(Address, Word)
	RAMAddressFetchUint32(Address) uint32
	RAMAddressWriteUint32(Address, uint32)
	RAMAddressFetchInt32(Address) int32
	RAMAddressWriteInt32(Address, int32)
	RAMAddressFetchBool(Address) bool
	RAMAddressWriteBool(Address, bool)
	RAMFrameFetch(FrameNumber) Frame
	RAMFrameWrite(FrameNumber, Frame)

	// Working with Disk
	DiskFrameFetch(FrameNumber) Frame
	DiskFrameWrite(FrameNumber, Frame)
}

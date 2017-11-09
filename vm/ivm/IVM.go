package ivm

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

	// Need the same interfaces used to get to RAM
	IRAM
}

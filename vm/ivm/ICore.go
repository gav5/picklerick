package ivm

// NumCoreRegisters is the number of registers in a given core.
const NumCoreRegisters = 16

// ICore is the kernel interface into a virtual machine core.
// (this is because vm uses the kernel, so we have to avoid circular dependencies)
// (this also helps regulate the access of the vm by the kernel to avoid abuse)
type ICore interface {

	// ProgramCounter manager
	ProgramCounter() Address
	SetProgramCounter(Address)

	// Program Halting
	Halt()
	Reset()

	// Managing Registers as Word
	RegisterWord(RegisterDesignation) Word
	SetRegisterWord(RegisterDesignation, Word)

	// Managing Registers as uint32
	RegisterUint32(RegisterDesignation) uint32
	SetRegisterUint32(RegisterDesignation, uint32)

	// Managing Registers as int32
	RegisterInt32(RegisterDesignation) int32
	SetRegisterInt32(RegisterDesignation, int32)

	// Managing Registers as bool
	RegisterBool(RegisterDesignation) bool
	SetRegisterBool(RegisterDesignation, bool)

	// Used for paging
	PagingProxy() PagingProxy
}

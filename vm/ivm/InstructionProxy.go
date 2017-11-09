package ivm

// InstructionProxy is used by an instruction to execute VM tasks.
type InstructionProxy struct {
	core ICore
	ram  IRAM
}

// MakeInstructionProxy makes an InstructionProxy instance
func MakeInstructionProxy(core ICore, ram IRAM) InstructionProxy {
	return InstructionProxy{
		core: core,
		ram:  ram,
	}
}

// ProgramCounter returns the value of the program counter.
func (ip InstructionProxy) ProgramCounter() Address {
	return ip.core.ProgramCounter()
}

// SetProgramCounter sets the program counter to the indicated value.
func (ip InstructionProxy) SetProgramCounter(val Address) {
	ip.core.SetProgramCounter(val)
}

// Halt halts program execution.
func (ip InstructionProxy) Halt() {
	ip.core.Halt()
}

// RegisterWord returns the given register word value.
func (ip InstructionProxy) RegisterWord(regNum RegisterDesignation) Word {
	return ip.core.RegisterWord(regNum)
}

// SetRegisterWord sets the given register to the given word value.
func (ip InstructionProxy) SetRegisterWord(regNum RegisterDesignation, val Word) {
	ip.core.SetRegisterWord(regNum, val)
}

// RegisterUint32 returns the given register uint32 value.
func (ip InstructionProxy) RegisterUint32(regNum RegisterDesignation) uint32 {
	return ip.core.RegisterUint32(regNum)
}

// SetRegisterUint32 sets the given register to the given uint32 value.
func (ip InstructionProxy) SetRegisterUint32(regNum RegisterDesignation, val uint32) {
	ip.core.SetRegisterUint32(regNum, val)
}

// RegisterInt32 returns the given register int32 value.
func (ip InstructionProxy) RegisterInt32(regNum RegisterDesignation) int32 {
	return ip.core.RegisterInt32(regNum)
}

// SetRegisterInt32 sets the given register to the given int32 value.
func (ip InstructionProxy) SetRegisterInt32(regNum RegisterDesignation, val int32) {
	ip.core.SetRegisterInt32(regNum, val)
}

// RegisterBool returns the given register bool value.
func (ip InstructionProxy) RegisterBool(regNum RegisterDesignation) bool {
	return ip.core.RegisterBool(regNum)
}

// SetRegisterBool sets the given register to the given bool value.
func (ip InstructionProxy) SetRegisterBool(regNum RegisterDesignation, val bool) {
	ip.core.SetRegisterBool(regNum, val)
}

// AddressFetchWord returns the word value at the given address.
func (ip InstructionProxy) AddressFetchWord(addr Address) Word {
	return ip.ram.AddressFetchWord(addr)
}

// AddressWriteWord writes the given word value to the given address.
func (ip InstructionProxy) AddressWriteWord(addr Address, val Word) {
	ip.ram.AddressWriteWord(addr, val)
}

// AddressFetchUint32 returns the uint32 value at the given address.
func (ip InstructionProxy) AddressFetchUint32(addr Address) uint32 {
	return ip.ram.AddressFetchUint32(addr)
}

// AddressWriteUint32 writes the given uint32 value to the given address.
func (ip InstructionProxy) AddressWriteUint32(addr Address, val uint32) {
	ip.ram.AddressWriteUint32(addr, val)
}

// AddressFetchInt32 returns the int32 value at the given address.
func (ip InstructionProxy) AddressFetchInt32(addr Address) int32 {
	return ip.ram.AddressFetchInt32(addr)
}

// AddressWriteInt32 writes the given int32 value to the given address.
func (ip InstructionProxy) AddressWriteInt32(addr Address, val int32) {
	ip.ram.AddressWriteInt32(addr, val)
}

// AddressFetchBool returns the bool value at the given address.
func (ip InstructionProxy) AddressFetchBool(addr Address) bool {
	return ip.ram.AddressFetchBool(addr)
}

// AddressWriteBool writes the given bool value to the given address.
func (ip InstructionProxy) AddressWriteBool(addr Address, val bool) {
	ip.ram.AddressWriteBool(addr, val)
}

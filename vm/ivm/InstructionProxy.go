package ivm

// InstructionProxy is used by an instruction to execute VM tasks.
type InstructionProxy struct {
	// core ICore
	// ram  IRAM
	// pagingProxy PagingProxy
	state *State
}

// MakeInstructionProxy makes an InstructionProxy instance
func MakeInstructionProxy(state *State) InstructionProxy {
	return InstructionProxy{state}
}

// ProgramCounter returns the value of the program counter.
func (ip InstructionProxy) ProgramCounter() Address {
	return ip.state.ProgramCounter
}

// SetProgramCounter sets the program counter to the indicated value.
func (ip InstructionProxy) SetProgramCounter(val Address) {
	ip.state.ProgramCounter = val
}

// Halt halts program execution.
func (ip InstructionProxy) Halt() {
	ip.state.Halt = true
}

// RegisterWord returns the given register word value.
func (ip InstructionProxy) RegisterWord(regNum RegisterDesignation) Word {
	return ip.state.RegisterWord(regNum)
}

// SetRegisterWord sets the given register to the given word value.
func (ip InstructionProxy) SetRegisterWord(regNum RegisterDesignation, val Word) {
	ip.state.SetRegisterWord(regNum, val)
}

// RegisterUint32 returns the given register uint32 value.
func (ip InstructionProxy) RegisterUint32(regNum RegisterDesignation) uint32 {
	return ip.state.RegisterUint32(regNum)
}

// SetRegisterUint32 sets the given register to the given uint32 value.
func (ip InstructionProxy) SetRegisterUint32(regNum RegisterDesignation, val uint32) {
	ip.state.SetRegisterUint32(regNum, val)
}

// RegisterInt32 returns the given register int32 value.
func (ip InstructionProxy) RegisterInt32(regNum RegisterDesignation) int32 {
	return ip.state.RegisterInt32(regNum)
}

// SetRegisterInt32 sets the given register to the given int32 value.
func (ip InstructionProxy) SetRegisterInt32(regNum RegisterDesignation, val int32) {
	ip.state.SetRegisterInt32(regNum, val)
}

// RegisterBool returns the given register bool value.
func (ip InstructionProxy) RegisterBool(regNum RegisterDesignation) bool {
	return ip.state.RegisterBool(regNum)
}

// SetRegisterBool sets the given register to the given bool value.
func (ip InstructionProxy) SetRegisterBool(regNum RegisterDesignation, val bool) {
	ip.state.SetRegisterBool(regNum, val)
}

func (ip InstructionProxy) translateAddress(addr Address) (FrameNumber, int) {
	return FrameNumber(addr/FrameSize), int(addr % FrameSize)
}

// AddressFetchWord returns the word value at the given address.
func (ip InstructionProxy) AddressFetchWord(addr Address) Word {
	frameNum, frameIndex := ip.translateAddress(addr)
	frame, ok := ip.state.Caches[frameNum]
	if !ok {
		// this is a fault, so we should add it
		ip.state.Faults.Set(frameNum)
		// return a blank value (it will produce garbage anyway)
		return 0x00000000
	}
	return frame[frameIndex]
}

// AddressWriteWord writes the given word value to the given address.
func (ip InstructionProxy) AddressWriteWord(addr Address, val Word) {
	frameNum, frameIndex := ip.translateAddress(addr)
	frame, ok := ip.state.Caches[frameNum]
	if !ok {
		// this is a fault, so we should add it and early exit
		ip.state.Faults.Set(frameNum)
		return
	}
	frame[frameIndex] = val
	ip.state.Caches[frameNum] = frame
}

// AddressFetchUint32 returns the uint32 value at the given address.
func (ip InstructionProxy) AddressFetchUint32(addr Address) uint32 {
	return ip.AddressFetchWord(addr).Uint32()
}

// AddressWriteUint32 writes the given uint32 value to the given address.
func (ip InstructionProxy) AddressWriteUint32(addr Address, val uint32) {
	ip.AddressWriteWord(addr, WordFromUint32(val))
}

// AddressFetchInt32 returns the int32 value at the given address.
func (ip InstructionProxy) AddressFetchInt32(addr Address) int32 {
	return ip.AddressFetchWord(addr).Int32()
}

// AddressWriteInt32 writes the given int32 value to the given address.
func (ip InstructionProxy) AddressWriteInt32(addr Address, val int32) {
	ip.AddressWriteWord(addr, WordFromInt32(val))
}

// AddressFetchBool returns the bool value at the given address.
func (ip InstructionProxy) AddressFetchBool(addr Address) bool {
	return ip.AddressFetchWord(addr).Bool()
}

// AddressWriteBool writes the given bool value to the given address.
func (ip InstructionProxy) AddressWriteBool(addr Address, val bool) {
	ip.AddressWriteWord(addr, WordFromBool(val))
}

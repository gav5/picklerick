package vm

import "./ivm"

// NumCores is the number of cores in the virtual machine.
const NumCores = 4

// VM is the virtual computer system.
type VM struct {
	Clock Clock
	Cores [NumCores]Core
	RAM   RAM
	Disk  Disk
}

func (vm VM) instructionProxy(corenum uint8) ivm.InstructionProxy {
	return ivm.MakeInstructionProxy(&vm.Cores[corenum], &vm.RAM)
}

// ProgramCounter returns the value of the program counter (PC).
func (vm VM) ProgramCounter(corenum uint8) ivm.Address {
	return vm.Cores[corenum].ProgramCounter()
}

// SetProgramCounter sets the value of the program counter to the given value.
func (vm *VM) SetProgramCounter(corenum uint8, addr ivm.Address) {
	vm.Cores[corenum].SetProgramCounter(addr)
}

// Halt halts current program execution.
func (vm *VM) Halt(corenum uint8) {
	vm.Cores[corenum].Halt()
}

// ResetCore resets the state to start over (undoing a halt)
func (vm *VM) ResetCore(corenum uint8) {
	vm.Cores[corenum].Reset()
}

// RegisterWord returns the given register word value.
func (vm VM) RegisterWord(corenum uint8, regNum ivm.RegisterDesignation) ivm.Word {
	return vm.Cores[corenum].RegisterWord(regNum)
}

// SetRegisterWord sets the given register to the given word value.
func (vm *VM) SetRegisterWord(corenum uint8, regNum ivm.RegisterDesignation, val ivm.Word) {
	vm.Cores[corenum].SetRegisterWord(regNum, val)
}

// RegisterUint32 returns the given register uint32 value.
func (vm VM) RegisterUint32(corenum uint8, regNum ivm.RegisterDesignation) uint32 {
	return vm.Cores[corenum].RegisterUint32(regNum)
}

// SetRegisterUint32 sets the given register to the given uint32 value.
func (vm *VM) SetRegisterUint32(corenum uint8, regNum ivm.RegisterDesignation, val uint32) {
	vm.Cores[corenum].SetRegisterUint32(regNum, val)
}

// RegisterInt32 returns the given register int32 value.
func (vm VM) RegisterInt32(corenum uint8, regNum ivm.RegisterDesignation) int32 {
	return vm.Cores[corenum].RegisterInt32(regNum)
}

// SetRegisterInt32 sets the given register to the given int32 value.
func (vm *VM) SetRegisterInt32(corenum uint8, regNum ivm.RegisterDesignation, val int32) {
	vm.Cores[corenum].SetRegisterInt32(regNum, val)
}

// RegisterBool returns the given register bool value.
func (vm VM) RegisterBool(corenum uint8, regNum ivm.RegisterDesignation) bool {
	return vm.Cores[corenum].RegisterBool(regNum)
}

// SetRegisterBool sets the given register to the given bool value.
func (vm *VM) SetRegisterBool(corenum uint8, regNum ivm.RegisterDesignation, val bool) {
	vm.Cores[corenum].SetRegisterBool(regNum, val)
}

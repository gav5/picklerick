package vm

import (
	"../disp"
	"./ivm"
)

// NumCores is the number of cores in the virtual machine.
const NumCores = 4

// VM is the virtual computer system.
type VM struct {
	Clock Clock
	Cores [NumCores]Core
	RAM   RAM
	Disk  Disk
}

const iter = 1000000

// MakeVM makes a new virtual machine.
func MakeVM() VM {
	return VM{
		Clock: 0x00000000,
		Cores: [NumCores]Core{
			MakeCore(0),
			MakeCore(1),
			MakeCore(2),
			MakeCore(3),
		},
		RAM:  MakeRAM(),
		Disk: Disk{},
	}
}

// Run runs the virtual machine.
func (vm *VM) Run(progress chan disp.Progress) {
	go func() {
		coreCh := make([]chan disp.Progress, NumCores)
		// run each core in its own goroutine
		for i, core := range vm.Cores {
			core.Run(coreCh[i], vm)
		}
		for {
		}
		// for i := 0; i < iter/2; i++ {
		// 	done <- disp.Progress{"Doing Thing (Part 1)", float32(i) / iter}
		// }
		// for i := iter / 2; i < iter; i++ {
		// 	done <- disp.Progress{"Doing Thing (Part 2)", float32(i) / iter}
		// }
		// done <- disp.Progress{"Done!", 1.0}
	}()
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

// RAMAddressFetchWord returns the word value at the given address.
func (vm VM) RAMAddressFetchWord(addr ivm.Address) ivm.Word {
	return vm.RAM.AddressFetchWord(addr)
}

// RAMAddressWriteWord writes the given word value to the given address.
func (vm *VM) RAMAddressWriteWord(addr ivm.Address, val ivm.Word) {
	vm.RAM.AddressWriteWord(addr, val)
}

// RAMAddressFetchUint32 returns the uint32 value at the given address.
func (vm VM) RAMAddressFetchUint32(addr ivm.Address) uint32 {
	return vm.RAM.AddressFetchUint32(addr)
}

// RAMAddressWriteUint32 writes the given uint32 value to the given address.
func (vm *VM) RAMAddressWriteUint32(addr ivm.Address, val uint32) {
	vm.RAM.AddressWriteUint32(addr, val)
}

// RAMAddressFetchInt32 returns the int32 value at the given address.
func (vm VM) RAMAddressFetchInt32(addr ivm.Address) int32 {
	return vm.RAM.AddressFetchInt32(addr)
}

// RAMAddressWriteInt32 writes the given int32 value to the given address.
func (vm *VM) RAMAddressWriteInt32(addr ivm.Address, val int32) {
	vm.RAM.AddressWriteInt32(addr, val)
}

// RAMAddressFetchBool returns the bool value at the given address.
func (vm VM) RAMAddressFetchBool(addr ivm.Address) bool {
	return vm.RAM.AddressFetchBool(addr)
}

// RAMAddressWriteBool writes the given bool value to the given address.
func (vm *VM) RAMAddressWriteBool(addr ivm.Address, val bool) {
	vm.RAM.AddressWriteBool(addr, val)
}

// RAMFrameFetch fetches the frame with the given frame number.
func (vm VM) RAMFrameFetch(frameNum ivm.FrameNumber) ivm.Frame {
	return vm.RAM.FrameFetch(frameNum)
}

// RAMFrameWrite writes the frame at the given frame number.
func (vm *VM) RAMFrameWrite(frameNum ivm.FrameNumber, frame ivm.Frame) {
	vm.RAM.FrameWrite(frameNum, frame)
}

// DiskFrameFetch fetches the frame with the given frame number.
func (vm VM) DiskFrameFetch(frameNum ivm.FrameNumber) ivm.Frame {
	return vm.Disk.FrameFetch(frameNum)
}

// DiskFrameWrite writes the frame at the given frame number.
func (vm *VM) DiskFrameWrite(frameNum ivm.FrameNumber, frame ivm.Frame) {
	vm.Disk.FrameWrite(frameNum, frame)
}

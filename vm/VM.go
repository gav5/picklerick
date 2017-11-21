package vm

import (
	"../disp"
	"./ivm"
	"./core"
	"../kernel"
	"../config"
	"log"
	"sync"
	"time"
)

// VM is the virtual computer system.
type VM struct {
	Clock Clock
	Cores [ivm.NumCores]*core.Core
	RAM   RAM
	Disk  Disk
	osKernel *kernel.Kernel
	reporter disp.ProgressReporter
	receiver disp.ProgressReceiver
}

const maxCount = 100

// New makes a new virtual machine.
func New(c config.Config) (*VM, error) {
	progress := make(chan disp.Progress)
	vm := &VM{
		Clock: 0x00000000,
		Cores: [ivm.NumCores]*core.Core{},
		RAM:  MakeRAM(),
		Disk: Disk{},
		reporter: disp.ProgressReporter(progress),
		receiver: disp.ProgressReceiver(progress),
	}
	// build each CPU Core
	for coreNum := uint8(0); coreNum < ivm.NumCores; coreNum++ {
		vm.Cores[coreNum] = core.New(coreNum)
	}
	// setup and configure the kernel
	var err error
	vm.osKernel, err = kernel.New(vm, c)
	if err != nil {
		return vm, err
	}
	return vm, nil
}

// Run runs the virtual machine.
func (vm *VM) Run() {
	for vm.Clock = 0x00000000; vm.Clock < maxCount; vm.Clock++ {
		log.Printf("[VM:%d] Tick!\n", vm.Clock)
		var wg sync.WaitGroup
		for coreNum, c := range vm.Cores {
			wg.Add(1)
			go func(c *core.Core, clock Clock) {
				defer wg.Done()
				log.Printf("[VM:%d] Sending context to core #%d...\n", clock, coreNum)
				c.CurrentContext = &core.Context{
					VM: vm,
					StartPC: 0x00000000,
				}
				log.Printf("[VM:%d] Running core #%d...\n", clock, coreNum)
				res := c.Call()
				// responseChan <- res
				if res.Error != nil {
					log.Printf(
						"[VM:%d] Error Running core #%d: %v\n",
						clock, coreNum, res.Error,
					)
					return
				}
				if res.Halted {
					log.Printf("[VM:%d] Core #%d has been HALTED\n", clock, coreNum)
					return
				}
			}(c, vm.Clock)
		}
		wg.Wait()

		// check if it's time to be done yet
		activeCores := ivm.NumCores
		for _, c := range vm.Cores {
			if c.ShouldHalt {
				activeCores--
			}
		}
		log.Printf(
			"[VM:%d] %d/%d active cores\n",
			vm.Clock, activeCores, ivm.NumCores,
		)
		if activeCores == 0 {
			break
		}
		time.Sleep(100)
	}
}

// InstructionProxy makes an instruction proxy for the given core
func (vm VM) InstructionProxy(c ivm.ICore) ivm.InstructionProxy {
	return ivm.MakeInstructionProxy(c, &vm.RAM)
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

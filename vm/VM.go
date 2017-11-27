package vm

import (
	"../disp"
	"./ivm"
	"./core"
	"../kernel"
	"../config"
	"../kernel/process"
	"log"
	"sync"
	"io"
	"fmt"
)

// VM is the virtual computer system.
type VM struct {
	Clock Clock
	Cores [ivm.NumCores]core.Core
	RAM   RAM
	Disk  Disk
	osKernel *kernel.Kernel
	reporter disp.ProgressReporter
	receiver disp.ProgressReceiver
	maxCycles uint
}

// New makes a new virtual machine.
func New(c config.Config) (*VM, error) {
	progress := make(chan disp.Progress)
	vm := &VM{
		Clock: 0x00000000,
		Cores: core.MakeArray(),
		RAM:  MakeRAM(),
		Disk: Disk{},
		reporter: disp.ProgressReporter(progress),
		receiver: disp.ProgressReceiver(progress),
		maxCycles: c.MaxCycles,
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
func (vm *VM) Run() error {
	for vm.Clock = 0x00000000; vm.Clock < Clock(vm.maxCycles); vm.Clock++ {
		// run a virtual machine cycle (and check for errors)
		if err := vm.runCycle(); err != nil {
			return err
		}
		// handle if there's nothing left to do (according to the Kernel)
		// (obviously when this is the case we should break out of the loop)
		if vm.osKernel.IsDone() {
			log.Printf("%s Kernel is DONE: breaking!\n", vm.callsign())
			return nil
		}
	}
	log.Printf("[VM:XXX] reached end of cycle limit: %d\n", vm.maxCycles)
	return nil
}

func (vm *VM) Tick() {
	vm.tick()
}

func (vm *VM) runCycle() error {

	// ensure the calls for Tick and Tock to the Kernel
	vm.tick()
	defer vm.tock()

	// setup and call each core in sequence
	// (wait for each to complete after that)
	var wg sync.WaitGroup
	for i := range vm.Cores {
		wg.Add(1)
		vm.setupCore(&vm.Cores[i])
		go func(c *core.Core){
			defer wg.Done()
			c.Call()
		}(&vm.Cores[i])
	}

	// after all cores are done, handle what they did
	wg.Wait()
	for i := range vm.Cores {
		err := vm.handleCore(&vm.Cores[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm VM) tick() {
	// log.Printf("%s Tick!\n", vm.callsign())
	vm.osKernel.Tick()
}

func (vm VM) tock() {
	// log.Printf("%s Tock!\n", vm.callsign())
	err := vm.osKernel.Tock()
	if err != nil {
		log.Printf("%s Tock error: %v\n", vm.callsign(), err)
		panic("Tock error: " + err.Error())
	}
}

// setupCore sets up the core to run some process.
// If necessary, it will get a new process from the Kernel.
func (vm VM) setupCore(c *core.Core) {
	if c.Process.IsSleep() {
		// we need to give this core a process!
		// the Kernel will know which one to do next!
		c.Process = vm.osKernel.ProcessForCore(c.CoreNum)

		// callsign := vm.callsign()
		// log.Printf(
		// 	"%s Setting up Core #%d with process #%d\n",
		// 	callsign, c.CoreNum, c.Process.ProcessNumber,
		// )
		// log.Printf(
		// 	"%s Core #%d has page table: %v\n",
		// 	callsign, c.CoreNum, c.Process.RAMPageTable,
		// )
		// log.Printf(
		// 	"%s Core #%d has caches: %v\n",
		// 	callsign, c.CoreNum, c.Process.State.Caches.Slice(),
		// )
	}
	c.Next = c.Process.State.Next()
}

// handleCore manages the final state from the execution of a core instruction.
// This unpacks that information and passes it to the Kernel.
func (vm VM) handleCore(c *core.Core) error {
	callsign := vm.callsign()
	// log.Printf("%s Handling Core #%d\n", callsign, c.CoreNum)

	if c.Process.IsSleep() {
		// there was no process, so an early exit is in order
		// (this is becasue the core was sleeping this cycle)
		return nil
	}

	if c.Next.Error != nil {
		// an error occured with the instruction execution
		log.Printf(
			"%s process %d threw an ERROR: %v\n",
			callsign, c.Process.ProcessNumber, c.Next.Error,
		)
		// stop the process and declare it a failure
		// (this should essentially be treated the same as a halt)
		c.Process.State = c.Process.State.Apply(c.Next)
		vm.osKernel.CompleteProcess(&c.Process)
		vm.osKernel.UpdateProcess(c.Process)
		// sine this is done, the process should be cleared
		// (this sends the message to later fill it if possible)
		c.Process = process.Sleep()

		// nil is returned here because there's nothing wrong with the VM
		// (it will just go to the next process like nothing ever happened)
		return nil
	}

	if c.Next.Halt {
		log.Printf(
			"%s process %d completed via HALT\n",
			callsign, c.Process.ProcessNumber,
		)
		// the core said to halt, so the process is now done!
		c.Process.State = c.Process.State.Apply(c.Next)
		vm.osKernel.CompleteProcess(&c.Process)
		vm.osKernel.UpdateProcess(c.Process)
		// since this is done, the process should be cleared
		// (this sends the message to later fill it if possible)
		c.Process = process.Sleep()
	} else if len(c.Next.Faults) > 0 {
		// looks like there were faults
		// (something was accessed that wasn't there)
		log.Printf(
			"%s process %d faulted: %v\n",
			callsign, c.Process.ProcessNumber, c.Next.Faults,
		)
		c.Process.Status = process.Wait
		// ensure the faults persist (and nothing else)
		c.Process.State.Faults = c.Next.Faults.Copy()
		vm.osKernel.UpdateProcess(c.Process)
		c.Process = process.Sleep()
	} else {
		// this was actually successful, so apply next so it's the actual state
		// log.Printf("%s applying next state\n", callsign)
		c.Process.State = c.Process.State.Apply(c.Next)
	}

	return nil
}

func (vm VM) callsign() string {
	return fmt.Sprintf("[VM:%d]", vm.Clock)
}

// FprintProcessTable prints the process table to the given writer.
func (vm VM) FprintProcessTable(w io.Writer) error {
	return vm.osKernel.FprintProcessTable(w)
}

// ProgramCounter returns the value of the program counter (PC).
// func (vm VM) ProgramCounter(corenum uint8) ivm.Address {
// 	return vm.Cores[corenum].Process.State.ProgramCounter
// }

// SetProgramCounter sets the value of the program counter to the given value.
// func (vm *VM) SetProgramCounter(corenum uint8, addr ivm.Address) {
// 	vm.Cores[corenum].Process.State.ProgramCounter = addr
// }

// Halt halts current program execution.
// func (vm *VM) Halt(corenum uint8) {
// 	vm.Cores[corenum].Halt()
// }

// ResetCore resets the state to start over (undoing a halt)
// func (vm *VM) ResetCore(corenum uint8) {
// 	vm.Cores[corenum].Reset()
// }

// RegisterWord returns the given register word value.
// func (vm VM) RegisterWord(corenum uint8, regNum ivm.RegisterDesignation) ivm.Word {
// 	return vm.Cores[corenum].RegisterWord(regNum)
// }

// SetRegisterWord sets the given register to the given word value.
// func (vm *VM) SetRegisterWord(corenum uint8, regNum ivm.RegisterDesignation, val ivm.Word) {
// 	vm.Cores[corenum].SetRegisterWord(regNum, val)
// }

// RegisterUint32 returns the given register uint32 value.
// func (vm VM) RegisterUint32(corenum uint8, regNum ivm.RegisterDesignation) uint32 {
// 	return vm.Cores[corenum].RegisterUint32(regNum)
// }

// SetRegisterUint32 sets the given register to the given uint32 value.
// func (vm *VM) SetRegisterUint32(corenum uint8, regNum ivm.RegisterDesignation, val uint32) {
// 	vm.Cores[corenum].SetRegisterUint32(regNum, val)
// }

// RegisterInt32 returns the given register int32 value.
// func (vm VM) RegisterInt32(corenum uint8, regNum ivm.RegisterDesignation) int32 {
// 	return vm.Cores[corenum].RegisterInt32(regNum)
// }

// SetRegisterInt32 sets the given register to the given int32 value.
// func (vm *VM) SetRegisterInt32(corenum uint8, regNum ivm.RegisterDesignation, val int32) {
// 	vm.Cores[corenum].SetRegisterInt32(regNum, val)
// }

// RegisterBool returns the given register bool value.
// func (vm VM) RegisterBool(corenum uint8, regNum ivm.RegisterDesignation) bool {
// 	return vm.Cores[corenum].RegisterBool(regNum)
// }

// SetRegisterBool sets the given register to the given bool value.
// func (vm *VM) SetRegisterBool(corenum uint8, regNum ivm.RegisterDesignation, val bool) {
// 	vm.Cores[corenum].SetRegisterBool(regNum, val)
// }

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

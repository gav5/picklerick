package vm

import (
	"../disp"
	"./ivm"
	"./core"
	"../kernel"
	"../config"
	"../kernel/process"
	"../util/logger"
	"sync"
	"io"
	"fmt"
	"log"
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
	logger *log.Logger
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
		logger: logger.New("vm"),
	}
	// setup and configure the kernel
	var err error
	vm.osKernel, err = kernel.New(vm, c)
	if err != nil {
		return vm, err
	}
	// run one tick to get everything together
	vm.tick()
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
			vm.logger.Printf("Kernel is DONE: breaking!")
			return nil
		}
	}
	vm.logger.Printf("reached end of cycle limit: %d", vm.maxCycles)
	return nil
}

func (vm *VM) runCycle() error {
	vm.logger.SetPrefix(vm.loggingPrefix())

	// ensure the calls for Tick and Tock to the Kernel
	vm.tick()
	defer vm.tock()

	// setup each core in sequence first
	processNums := make([]uint8, ivm.NumCores)
	for i := range vm.Cores {
		vm.setupCore(&vm.Cores[i])
		// make sure this one hasn't been given out already
		processNumber := vm.Cores[i].Process.ProcessNumber
		for _, pn := range processNums {
			if processNumber == pn {
				// this one has already been handed out!
				// this should get reported
				return ProcessAllocationError{vm.Cores}
			}
		}
		processNums[i] = processNumber
	}
	vm.logger.Printf("Process Allocation: %v", processNums)

	// call each core in sequence
	// (wait for each to complete after that)
	var wg sync.WaitGroup
	for i := range vm.Cores {
		wg.Add(1)
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

// ProcessAllocationError is when the same process is allocated more than once.
type ProcessAllocationError struct {
	cores [ivm.NumCores]core.Core
}
func (err ProcessAllocationError) Error() string {
	processNums := [ivm.NumCores]uint8{}
	for i, c := range err.cores {
		processNums[i] = c.Process.ProcessNumber
	}
	return fmt.Sprintf(
		"the same process has been over-allocated: %v", processNums,
	)
}


func (vm VM) tick() {
	// logger.Printf("%s Tick!\n", vm.callsign())
	vm.osKernel.Tick()
}

func (vm VM) tock() {
	// logger.Printf("%s Tock!\n", vm.callsign())
	err := vm.osKernel.Tock()
	if err != nil {
		vm.logger.Printf("Tock error: %v", err)
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
	}
	c.Next = c.Process.State.Next()
}

// handleCore manages the final state from the execution of a core instruction.
// This unpacks that information and passes it to the Kernel.
func (vm VM) handleCore(c *core.Core) error {
	// logger.Printf("%s Handling Core #%d\n", callsign, c.CoreNum)

	if c.Process.IsSleep() {
		// there was no process, so an early exit is in order
		// (this is becasue the core was sleeping this cycle)
		return nil
	}

	if c.Next.Error != nil {
		// an error occured with the instruction execution
		vm.logger.Printf(
			"process %d threw an ERROR: %v\n",
			c.Process.ProcessNumber, c.Next.Error,
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
		vm.logger.Printf(
			"process %d completed via HALT",
			c.Process.ProcessNumber,
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
		vm.logger.Printf(
			"process %d faulted: %v\n",
			c.Process.ProcessNumber, c.Next.Faults,
		)
		c.Process.SetStatus(process.Wait)
		// ensure the faults persist (and nothing else)
		c.Process.State.Faults = c.Next.Faults.Copy()
		vm.osKernel.UpdateProcess(c.Process)
		c.Process = process.Sleep()
	} else {
		// this was actually successful, so apply next so it's the actual state
		// logger.Printf("%s applying next state\n", callsign)
		c.Process.State = c.Process.State.Apply(c.Next)
	}

	return nil
}

// FprintProcessTable prints the process table to the given writer.
func (vm VM) FprintProcessTable(w io.Writer) error {
	return vm.osKernel.FprintProcessTable(w)
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

func (vm VM) loggingPrefix() string {
	return fmt.Sprintf("vm:%d | ", vm.Clock)
}

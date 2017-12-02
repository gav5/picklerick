package process

import (
	"fmt"
	"log"

	"../../util/logger"
	"../../vm/ivm"
	"../page"
	"../program"
)

// Process is used to represent process execution
// also used to facilitate CPU reassignment during context switching
type Process struct {
	// CPUID is used to describe the CPU the process is being run on
	CPUID uint8
	// CodeSize indicates the size of the code for the given process
	CodeSize uint8
	// Registers contains the list of standard CPU registers
	// (these are manipulated by instructions for general computational purposes)
	Registers [ivm.NumCoreRegisters]ivm.Word
	// ProcessNumber is the number assigned to the process for tracking
	ProcessNumber uint8
	// RAMPageTable is used to track all RAM pages used by the process
	RAMPageTable page.Table
	// DiskPageTable is used to track all Disk pages used by the process
	DiskPageTable page.Table
	// Priority is used for sorting purposes
	Priority uint8
	// Metrics holds all the data about the process
	Metrics Metrics
	// Program holds the original program that generated the process
	Program program.Program
	// Footprint stores the number of frames/pages required to store this process
	Footprint int

	state   ivm.State
	status  Status
	isSleep bool
	logger  *log.Logger
}

// Make makes a Process from a given program
func Make(p program.Program) Process {
	return Process{
		CPUID:         0x0,
		CodeSize:      p.NumberOfWords,
		ProcessNumber: p.JobID,
		Priority:      p.PriorityNumber,
		RAMPageTable:  make(page.Table),
		DiskPageTable: make(page.Table),
		Metrics:       MakeMetrics(),
		Footprint:     4,
		Program:       p,
		state:         ivm.MakeState(),
		status:        New,
		isSleep:       false,
		logger:        logger.New(fmt.Sprintf("process%02d", p.JobID)),
	}
}

// Mock produces a fake process for testing.
func Mock(p program.Program) Process {
	return Process{
		CPUID:         0x0,
		CodeSize:      p.NumberOfWords,
		ProcessNumber: p.JobID,
		Priority:      p.PriorityNumber,
		RAMPageTable:  make(page.Table),
		DiskPageTable: make(page.Table),
		Metrics:       MakeMetrics(),
		Footprint:     4,
		Program:       p,
		state: ivm.State{
			ProgramCounter: 0x00,
			Halt:           false,
			Error:          nil,
			Registers:      [ivm.NumCoreRegisters]ivm.Word{},
			Faults:         map[ivm.FrameNumber]bool{},
			Caches:         ivm.FrameCacheArrayFromUint32Array(p.RAMRepresentation()),
		},
		status:  New,
		isSleep: false,
		logger:  logger.Dummy(),
	}
}

// Sleep makes a process that tells the CPU to sleep each time.
func Sleep() Process {
	return Process{
		CPUID:         0x0,
		CodeSize:      1,
		ProcessNumber: 0,
		Priority:      0,
		RAMPageTable:  make(page.Table),
		DiskPageTable: make(page.Table),
		Footprint:     0,
		Program:       program.Sleep(),
		state:         ivm.Sleep(),
		status:        Ready,
		isSleep:       true,
		logger:        logger.Dummy(),
	}
}

// Copy produces a duplicate process that is the same as the given one.
func (p Process) Copy() Process {
	return Process{
		CPUID:         p.CPUID,
		CodeSize:      p.CodeSize,
		ProcessNumber: p.ProcessNumber,
		Priority:      p.Priority,
		RAMPageTable:  p.RAMPageTable.Copy(),
		DiskPageTable: p.DiskPageTable.Copy(),
		Footprint:     p.Footprint,
		Program:       p.Program.Copy(),
		state:         p.state,
		status:        p.status,
		isSleep:       p.isSleep,
		logger:        p.logger,
	}
}

// State contains the operational state of the CPU
func (p Process) State() ivm.State {
	return p.state
}

// SetState mutates the state of the process to a new value.
func (p *Process) SetState(val ivm.State) {
	p.logger.Printf("set state to %#v", val)
	p.logger.Printf("state currently %#v", p.state)

	(*p).state = val
	p.logger.Printf("state is now %#v", p.state)
}

// Status describes the current status of the process
// (ex: if it is running, waiting, etc)
func (p Process) Status() Status {
	return p.status
}

// SetStatus mutates the status to a new value.
func (p *Process) SetStatus(val Status) {
	p.logger.Printf(
		"set status from %v to %v",
		p.status, val,
	)
	if !validateTransition(p.status, val) {
		p.logger.Panicf(
			"cannot transition from %v to %v",
			p.status, val,
		)
	}
	old := p.status
	(*p).status = val
	(*p).Metrics.ReactToStatus(p.status, val)
	p.logger.Printf("status now %v (was %v)", p.status, old)
}

// InputBufferOffset returns the offset to use to get the input buffer.
func (p Process) InputBufferOffset() int {
	return int(p.Program.NumberOfWords)
}

// InputBufferRange returns the range of addresses used for the input buffer.
func (p Process) InputBufferRange() (ivm.Address, ivm.Address) {
	offset := p.InputBufferOffset()
	addr := ivm.Address(offset) * 4
	return addr, addr + ivm.Address((p.Program.InputBufferSize-1)*4)
}

// InputBuffer returns the input section of the process.
// This is for reporting use only (use caches for program execution)
// This feeds from disk, so it will only work after the whole system is complete
func (p Process) InputBuffer(v ivm.IVM) ([]ivm.Word, int) {
	return p.fetchFromResultBuffer(
		p.InputBufferOffset(),
		p.Program.InputBufferSize,
		v,
	)
}

// TempBufferOffset returns the offset to use to get the temp buffer.
func (p Process) TempBufferOffset() int {
	return int(p.Program.NumberOfWords + p.Program.InputBufferSize + p.Program.OutputBufferSize)
}

// TempBufferRange returns the range of addresses used for the temp buffer.
func (p Process) TempBufferRange() (ivm.Address, ivm.Address) {
	return p.resultBufferRange(p.TempBufferOffset(), p.Program.TempBufferSize)
}

// TempBuffer returns the temp section of the process.
// This is for reporting use only (use caches for program execution)
func (p Process) TempBuffer(v ivm.IVM) ([]ivm.Word, int) {
	return p.fetchFromResultBuffer(
		p.TempBufferOffset(),
		p.Program.TempBufferSize,
		v,
	)
}

// OutputBufferOffset returns the offset to use to get the output buffer.
func (p Process) OutputBufferOffset() int {
	return int(p.Program.NumberOfWords + p.Program.InputBufferSize)
}

// OutputBufferRange returns the range of addresses used for the output buffer.
func (p Process) OutputBufferRange() (ivm.Address, ivm.Address) {
	return p.resultBufferRange(
		p.OutputBufferOffset(),
		p.Program.OutputBufferSize,
	)
}

// OutputBuffer returns the output section of the process.
// This is for reporting use only (use caches for program execution)
func (p Process) OutputBuffer(v ivm.IVM) ([]ivm.Word, int) {
	return p.fetchFromResultBuffer(
		p.OutputBufferOffset(),
		p.Program.OutputBufferSize,
		v,
	)
}

func (p Process) resultBufferRange(offset int, size uint8) (ivm.Address, ivm.Address) {
	addr := ivm.AddressForIndex(offset)
	return addr, addr + ivm.AddressForIndex(int(size)-1)
}

func (p Process) fetchFromResultBuffer(offset int, size uint8, v ivm.IVM) ([]ivm.Word, int) {
	buf := make([]ivm.Word, size)
	dataOffset := offset - int(p.Program.NumberOfWords)
	for i := range buf {
		addr := ivm.AddressForIndex(offset + i)
		frameNumber, frameIndex, ok := p.DiskPageTable.PairForAddressSoft(addr)
		if ok {
			// get the value from disk (because it's been changed by the program)
			frame := v.DiskFrameFetch(frameNumber)
			buf[i] = frame[frameIndex]
		} else {
			frameNumber, frameIndex, ok = p.RAMPageTable.PairForAddressSoft(addr)
			if ok {
				// get the value from RAM (because it was never saved back)
				frame := v.RAMFrameFetch(frameNumber)
				buf[i] = frame[frameIndex]
			} else {
				// get the value from the original program (because it's never been touched by the program)
				buf[i] = ivm.Word(p.Program.DataBlock[dataOffset+i])
			}
		}
	}
	return buf, offset
}

// MakeArray makes an array of Processes from a given array of programs.
func MakeArray(progAry []program.Program) []Process {
	outary := make([]Process, len(progAry))
	for i, prog := range progAry {
		outary[i] = Make(prog)
	}
	return outary
}

// IsSleep returns if this is a sleep process or not
func (p Process) IsSleep() bool {
	return p.isSleep
}

// TableHeaders describes the headers used for the table of processes.
func TableHeaders() []string {
	return []string{
		"ProcessNumber",
		"Status",
		"Priority",
		"Size",
		"RAM Pages",
		"Disk Pages",
		"Completion Time (ns)",
		"Wait Time (ns)",
	}
}

// TableRow returns the corresponding row for the given process.
func (p Process) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", p.ProcessNumber),
		fmt.Sprintf("%v", p.Status()),
		fmt.Sprintf("%d", p.Priority),
		fmt.Sprintf("%d", p.CodeSize),
		fmt.Sprintf("%d", len(p.RAMPageTable)),
		fmt.Sprintf("%d", len(p.DiskPageTable)),
		fmt.Sprintf("%s", p.completionTimeLabel()),
		fmt.Sprintf("%d", p.Metrics.WaitTime.Value().Nanoseconds()),
	}
}

func (p Process) completionTimeLabel() string {
	if p.status != Terminated {
		return ""
	}
	return fmt.Sprintf("%d", p.Metrics.CompletionTime.Value().Nanoseconds())
}

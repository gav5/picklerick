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
	// ProgramCounter (PC) describes where the program is at in execution
	ProgramCounter ivm.Address
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
		CPUID:          0x0,
		ProgramCounter: 0x00,
		CodeSize:       p.NumberOfWords,
		ProcessNumber:  p.JobID,
		Priority:       p.PriorityNumber,
		RAMPageTable:   make(page.Table),
		DiskPageTable:  make(page.Table),
		Footprint:      4,
		Program:        p,
		state:          ivm.MakeState(),
		status:         New,
		isSleep:        false,
		logger:         logger.New(fmt.Sprintf("process%02d", p.JobID)),
	}
}

// Sleep makes a process that tells the CPU to sleep each time.
func Sleep() Process {
	return Process{
		CPUID:          0x0,
		ProgramCounter: 0x00,
		CodeSize:       1,
		ProcessNumber:  0,
		Priority:       0,
		RAMPageTable:   make(page.Table),
		DiskPageTable:  make(page.Table),
		Footprint:      0,
		Program:        program.Sleep(),
		state:          ivm.Sleep(),
		status:         Ready,
		isSleep:        true,
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
	(*p).Metrics.ReactToStatus(p.status, val)
	(*p).status = val
	p.logger.Printf("status now %v (was %v)", p.status, old)
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
	}
}

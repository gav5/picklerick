package process

import (
	"../page"
	"../program"
	"../../metric"
	"../../vm/ivm"
)

// Process is used to represent process execution
// also used to facilitate CPU reassignment during context switching
type Process struct {

	// CPUID is used to describe the CPU the process is being run on
	CPUID uint8

	// ProgramCounter (PC) describes where the program is at in execution
	ProgramCounter ivm.Address

	// State contains the operational state of the CPU
	State ivm.State

	// CodeSize indicates the size of the code for the given process
	CodeSize uint8

	// Registers contains the list of standard CPU registers
	// (these are manipulated by instructions for general computational purposes)
	Registers [ivm.NumCoreRegisters]ivm.Word

	// ProcessNumber is the number assigned to the process for tracking in the process table
	ProcessNumber uint8

	// RAMPageTable is used to track all RAM pages used by the process
	RAMPageTable page.Table

	// DiskPageTable is used to track all Disk pages used by the process
	DiskPageTable page.Table

	// Priority is used for sorting purposes
	Priority uint8

	Metrics struct {
		JobWaitTime       metric.StopwatchMetric
		JobCompletionTime metric.StopwatchMetric
		IOOperationCount  metric.CountMetricUint32
		RAMUse            metric.FractionalMetricUint32
		CacheUse          metric.FractionalMetricUint32
	}

	Program program.Program

	// Footprint stores the number of frames/pages required to store this process
	Footprint int

	// State describes the current status of the process
	// (ex: if it is running, waiting, etc)
	Status Status

	isSleep bool
}

// Make makes a Process from a given program
func Make(p program.Program) Process {
	return Process{
		CPUID:          0x0,
		ProgramCounter: 0x00,
		CodeSize:       p.NumberOfWords,
		ProcessNumber:  p.JobID,
		Priority: 			p.PriorityNumber,
		RAMPageTable:  	make(page.Table),
		DiskPageTable:	make(page.Table),
		Footprint:			4,
		Program: 				p,
		State:					ivm.MakeState(),
		Status:					New,
		isSleep:				false,
	}
}

// Sleep makes a process that tells the CPU to sleep each time.
func Sleep() Process {
	return Process{
		CPUID: 0x0,
		ProgramCounter: 0x00,
		CodeSize: 1,
		ProcessNumber: 0,
		Priority: 0,
		RAMPageTable: make(page.Table),
		DiskPageTable: make(page.Table),
		Footprint: 0,
		Program: program.Sleep(),
		State: ivm.Sleep(),
		Status: Ready,
		isSleep: true,
	}
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

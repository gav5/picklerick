package process

import (
	"../page"
	"../../prog"
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
	// State cpu.State

	// CodeSize indicates the size of the code for the given process
	CodeSize uint8

	// Registers contains the list of standard CPU registers
	// (these are manipulated by instructions for general computational purposes)
	Registers [ivm.NumCoreRegisters]ivm.Word

	// ProcessNumber is the number assigned to the process for tracking in the process table
	ProcessNumber uint8

	// PageTable is used to track all pages used by the process
	PageTable page.Table

	// Priority is used for sorting purposes
	Priority uint8

	Metrics struct {
		JobWaitTime       metric.StopwatchMetric
		JobCompletionTime metric.StopwatchMetric
		IOOperationCount  metric.CountMetricUint32
		RAMUse            metric.FractionalMetricUint32
		CacheUse          metric.FractionalMetricUint32
	}

	// Program describes the program the PCB is running
	// NOTE: this is a temporary measure to make this work!
	// Program prog.Program

	// schedule: any
	// accounts: any

	// Memories Memories

	// progeny: any
	// ptr: any
	// resources: any

	// Status describes the current status of the process
	// (ex: if it is running, waiting, etc)
	// Status Status

	// statusInfo: any
	// priority: any
}

// Make makes a Process from a given program and page table
func Make(program prog.Program, pageTable page.Table) Process {
	return Process{
		CPUID:          0x0,
		ProgramCounter: 0x00,
		CodeSize:       program.Job.NumberOfWords,
		ProcessNumber:  program.Job.ID,
		Priority: 			program.Job.PriorityNumber,
		PageTable:      pageTable,
	}
}

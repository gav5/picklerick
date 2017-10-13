package proc

import (
	"../cpu"
	"../reg"
)

// PCB (aka Process Control Block) used to represent process execution
// also used to facilitate CPU reassignment during context switching
type PCB struct {

	// CPUID is used to describe the CPU the process is being run on
	CPUID cpu.ID

	// ProgramCounter (PC) describes where the program is at in execution
	ProgramCounter reg.Storage

	// State contains the operational state of the CPU
	State cpu.State

	// CodeSize indicates the size of the code for the given process
	CodeSize uint8

	// Registers contains the list of standard CPU registers
	// (these are manipulated by instructions for general computational purposes)
	Registers reg.List

	// schedule: any
	// accounts: any

	Memories Memories

	// progeny: any
	// ptr: any
	// resources: any

	// Status describes the current status of the process
	// (ex: if it is running, waiting, etc)
	Status Status

	// statusInfo: any
	// priority: any
}

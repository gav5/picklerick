package main

// PCB (aka Process Control Block) used to represent process execution
// also used to facilitate CPU reassignment during context switching
type PCB struct {
	// cpuid: any
	// programCounter: any
	state CPUState
	// codeSize: any
	// registers: any
	// schedule: any
	// accounts: any
	// memories: any
	// progeny: any
	// ptr: any
	// resources: any
	// Determines the current status of the process
	// (ex: if it is running, waiting, etc)
	status ProcessStatus

	// statusInfo: any
	// priority: any
}

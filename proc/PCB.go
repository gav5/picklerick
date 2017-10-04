package proc

import "../cpu"

// PCB (aka Process Control Block) used to represent process execution
// also used to facilitate CPU reassignment during context switching
type PCB struct {
	// cpuid: any
	// programCounter: any
	state cpu.State
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
	status Status

	// statusInfo: any
	// priority: any
}

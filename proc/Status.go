package proc

// Status describes the status of a given process
type Status int

const (
	// New describes a fresh, newly-created process
	New Status = iota
	// Ready describes a process that is ready to do something
	Ready
	// Blocked describes a process that is currently held up waiting on something
	Blocked
	// Running describes a process in the middle of execution
	Running
)

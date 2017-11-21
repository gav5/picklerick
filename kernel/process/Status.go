package process

// Status indicates what a given process is doing at the moment.
type Status uint8

const (
  // Waiting means the process is waiting in the queue.
  Waiting Status = iota
  // Running means the process is running on a CPU Core.
  Running
  // Done means the process is done being run and should be killed.
  Done
  // Dead means the process has completed and is off the system.
  Dead
)

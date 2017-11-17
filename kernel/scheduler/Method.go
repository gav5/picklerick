package scheduler

import (
  // "sort"
  kernel ".."
)

// Method defines how to sort a list of processes.
type Method func(p1, p2 kernel.Process) bool

const (

  // SwitchFIFO describes the CLI switch used to indicate FIFO sorting
  SwitchFIFO = "fifo"

  // SwitchPriority describes the CLI switch used to indicate Priority sorting
  SwitchPriority = "priority"
)

// FIFO sorts processes using a First-In First-Out policy.
func FIFO(p1, p2 kernel.Process) bool {
  // false is returned here so nothing gets moved
  return false
}

// Priority sorts processes using the priority number of the process.
func Priority(p1, p2 kernel.Process) bool {
  return p1.Priority < p2.Priority
}

// MethodForSwitch returns a Method for the given switch arg.
func MethodForSwitch(sw string) Method {
  switch sw {
  case SwitchFIFO:
    return FIFO
  case SwitchPriority:
    return Priority
  default:
    // if nothing else provided, assume FIFO
    return FIFO
  }
}

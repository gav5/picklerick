package scheduler

import (
  // "sort"
  "../process"
)

// Method defines how to sort a list of processes.
type Method func(p1, p2 process.Process) bool

const (

  // SwitchFCFS describes the switch for first come first serve (FCFS)
  SwitchFCFS = "fcfs"

  // SwitchPriority describes the CLI switch used to indicate Priority sorting
  SwitchPriority = "priority"

  // SwitchSJF describes the CLI switch for shortest job first (SJF)
  SwitchSJF = "sjf"
)

// FCFS sorts processes using a First-Come-First-Serve policy.
func FCFS(p1, p2 process.Process) bool {
  // sort by the process number (low to high)
  return p1.ProcessNumber > p2.ProcessNumber
}

// Priority sorts processes using the priority number of the process.
func Priority(p1, p2 process.Process) bool {
  return p1.Priority < p2.Priority
}

// SJF sorts the processes using by the job's number of instructions.
func SJF(p1, p2 process.Process) bool {
  return p1.CodeSize > p2.CodeSize
}

// MethodForSwitch returns a Method for the given switch arg.
func MethodForSwitch(sw string) Method {
  switch sw {
  case SwitchFCFS:
    return FCFS
  case SwitchPriority:
    return Priority
  case SwitchSJF:
    return SJF
  default:
    // if nothing else provided, assume FCFS
    return FCFS
  }
}

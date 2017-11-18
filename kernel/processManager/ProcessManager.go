package processManager

import (
  "container/heap"

  "../process"
)

// TODO: make into a priority queue
// (look at https://golang.org/pkg/container/heap/#example__priorityQueue)

// SortMethod defines a function that can be used to sort processes.
type SortMethod func(p1, p2 process.Process) bool

// ProcessManager keeps track of system processes.
type ProcessManager struct {
  processList processList
}

// New creates a new process manager.
func New(sortMethod SortMethod) *ProcessManager {
  pm := &ProcessManager{
    processList: processList{
      base: []process.Process{},
      sortMethod: sortMethod,
    },
  }
  heap.Init(&pm.processList)
  return pm
}

// Add adds a process into the process manager.
func (pm ProcessManager) Add(p process.Process) {
  heap.Push(pm.processList, p)
}

// Find returns the process with the corresponding process number.
func (pm ProcessManager) Find(processNumber uint8) *process.Process {
  for _, p := range pm.processList.base {
    if p.ProcessNumber == processNumber {
      return &p
    }
  }
  return nil
}

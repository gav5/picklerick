package processManager

import (
  "container/heap"
  "log"

  "../process"
  "../../vm/ivm"
)

// TODO: make into a priority queue
// (look at https://golang.org/pkg/container/heap/#example__priorityQueue)

// SortMethod defines a function that can be used to sort processes.
type SortMethod func(p1, p2 process.Process) bool

// ProcessManager keeps track of system processes.
type ProcessManager struct {
  processList *processList
  completed []process.Process
}

// New creates a new process manager.
func New(sortMethod SortMethod) *ProcessManager {
  pm := &ProcessManager{
    processList: &processList{
      base: []process.Process{},
      sortMethod: sortMethod,
    },
  }
  heap.Init(pm.processList)
  return pm
}

// ProcessForCore returns the appropriate process for the given core.
func (pm ProcessManager) ProcessForCore(coreNum int) *process.Process {
  return &pm.processList.base[coreNum]
}

// ProcessesForQueue returns processes that should be queued.
func (pm ProcessManager) ProcessesForQueue() []*process.Process {
  procSlice := pm.processList.base[ivm.NumCores:]
  outary := make([]*process.Process, pm.processList.Len() - ivm.NumCores)
  for i := range outary {
    outary[i] = &procSlice[i]
  }
  return outary
}

// Reevaluate re-sorts the current process list and excludes unnecessary ones.
func (pm ProcessManager) Reevaluate() {
  for i, p := range pm.processList.base[:ivm.NumCores] {
    if p.Status == process.Done {
      log.Printf("process %d has now been completed!\n", p.ProcessNumber)
      pm.processList.base = append(
        pm.processList.base[:i],
        pm.processList.base[i+1:]...,
      )
    }
  }
}

// IsDone returns if the system is done yet.
func (pm ProcessManager) IsDone() bool {
  return pm.processList.Len() == 0
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

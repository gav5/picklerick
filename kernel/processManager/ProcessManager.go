package processManager

import "../process"

// TODO: make into a priority queue
// (look at https://golang.org/pkg/container/heap/#example__priorityQueue)

// SortMethod defines a function that can be used to sort processes.
type SortMethod func(p1, p2 process.Process) bool

// ProcessManager keeps track of system processes.
type ProcessManager struct {
  processList []process.Process
  sortMethod SortMethod
}

// New creates a new process manager.
func New(sortMethod SortMethod) *ProcessManager {
  return &ProcessManager{
    processList: []process.Process{},
    sortMethod: sortMethod,
  }
}

// Len is for the heap interface
func (pm ProcessManager) Len() int {
  return len(pm.processList)
}

// Less is for the heap interface
// (this is where sorting is done!!)
func (pm ProcessManager) Less(i, j int) bool {
  return pm.sortMethod(pm.processList[i], pm.processList[j])
}

// Swap is for the heap interface
func (pm ProcessManager) Swap(i, j int) {
  pm.processList[i], pm.processList[j] = pm.processList[j], pm.processList[i]
}

// Push is for the heap interface
func (pm *ProcessManager) Push(x interface{}) {
  pm.processList = append(pm.processList, x.(process.Process))
}

// Pop is for the heap interface
func (pm *ProcessManager) Pop() interface{} {
  old := pm.processList
  n := len(old)
  x := old[n-1]
  pm.processList = old[0:n-1]
  return x
}

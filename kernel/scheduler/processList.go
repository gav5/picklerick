package scheduler

import "../process"

type processList struct {
  base []process.Process
  sortMethod Method
}

func (pl processList) Len() int {
  return len(pl.base)
}

func (pl processList) Less(i, j int) bool {
  return pl.sortMethod(pl.base[i], pl.base[j])
}

// Swap is for the heap interface
func (pl processList) Swap(i, j int) {
  pl.base[i], pl.base[j] = pl.base[j], pl.base[i]
}

// Push is for the heap interface
func (pl *processList) Push(x interface{}) {
  (*pl).base = append(pl.base, x.(process.Process))
}

// Pop is for the heap interface
func (pl *processList) Pop() interface{} {
  old := pl.base
  n := len(old)
  x := old[n-1]
  (*pl).base = old[0:n-1]
  return x
}

// NOTE: need to load into ProcessManager first
// (then add elements to RAM, add remaining to disk)

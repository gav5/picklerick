package scheduler

import (
  "sort"
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

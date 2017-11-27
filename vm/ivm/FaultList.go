package ivm

import "fmt"

// FaultList is for keeping track of faults.
type FaultList map[FrameNumber]bool

// Check returns if the fault list has the given frame number.
func (fl FaultList) Check(fn FrameNumber) bool {
  return fl[fn]
}

// Set ensures the fault list has the given frame number.
func (fl *FaultList) Set(fn FrameNumber) {
  (*fl)[fn] = true
}

// Slice returns the list of faulted frame numbers as a slice.
func (fl FaultList) Slice() []FrameNumber {
  s := make([]FrameNumber, len(fl))
  i := 0
  for fn := range fl {
    s[i] = fn
    i++
  }
  return s
}

func (fl FaultList) String() string {
  return fmt.Sprintf("%v", fl.Slice())
}

// Copy creates a duplicate FaultList
func (fl FaultList) Copy() FaultList {
  c := make(FaultList)
  for fn, v := range fl {
    c[fn] = v
  }
  return c
}

package ivm

import "fmt"

// FrameNumber describes the number of a frame.
type FrameNumber int

func (fn FrameNumber) String() string {
  return fmt.Sprintf("%03X", int(fn))
}

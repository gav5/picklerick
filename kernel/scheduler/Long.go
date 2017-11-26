package scheduler

import (
  "fmt"

  "../process"
)

// LongTermQueueThreshhold describes how much RAM is the long-term queue.
const LongTermQueueThreshhold = 512

// Long performs long-term scheduling operations.
// This means adding any processes to RAM if possible.
func (sched Scheduler) Long() {
  // go through each process until the RAM is sufficiently filled
  // this is a tunable value, of course (since we want some spare space)
  sched.EachWhile(func (p *process.Process) bool {
    if sched.pm.AvailableRAM() >= LongTermQueueThreshhold {
      // the threshhold has been met (need to use the RAM for other stuff)
      // we should break out now and leave things as they are
      return false
    }
    // fill the RAM with another process's initial footprint
    if err := sched.pm.Load(p); err != nil {
      panic(fmt.Sprintf("failed to load process %v\n", *p))
      return false
    }
    return false
  })
}

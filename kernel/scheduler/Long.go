package scheduler

import (
  "fmt"

  "../process"
  "../../vm/ivm"
)

// Long performs long-term scheduling operations.
// This means adding any processes to RAM if possible.
func (sched Scheduler) Long() {
  // go through each process until the RAM is sufficiently filled
  // this is a tunable value, of course (since we want some spare space)
  leftoverSpace := int(ivm.RAMNumFrames - sched.longTermQueueSize)
  sched.logger.Printf(
    "[Long] Using queue of size %d",
    sched.longTermQueueSize,
  )
  sched.logger.Printf(
    "[Long] A cushion of %d will be left in RAM",
    leftoverSpace,
  )
  sched.EachWhile(func (p *process.Process) bool {
    avail := sched.pm.AvailableRAM()
    if avail < leftoverSpace {
      // the threshhold has been met (need to use the RAM for other stuff)
      // we should break out now and leave things as they are
      sched.logger.Printf(
        "[Long] Stopping at RAM %d/%d", avail, ivm.RAMNumFrames,
      )
      return false
    }
    // fill the RAM with another process's initial footprint
    if err := sched.pm.Load(p); err != nil {
      msg := fmt.Sprintf("failed to load process %v", *p)
      sched.logger.Printf(
        "[Long] RAM is at %d/%d",
        sched.pm.AvailableRAM(),
        ivm.RAMNumFrames,
      )
      sched.logger.Printf("[Long] %s\n", msg)
      panic(msg)
      return false
    }
    sched.logger.Printf(
      "[Long] process %d is loaded (RAM is now %d/%d)\n",
      p.ProcessNumber, sched.pm.AvailableRAM(), ivm.RAMNumFrames,
    )
    return true
  })
}

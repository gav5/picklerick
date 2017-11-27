package scheduler

import (
  "fmt"
  "log"

  "../process"
  "../../vm/ivm"
)

// Long performs long-term scheduling operations.
// This means adding any processes to RAM if possible.
func (sched Scheduler) Long() {
  // go through each process until the RAM is sufficiently filled
  // this is a tunable value, of course (since we want some spare space)
  leftoverSpace := int(ivm.RAMNumFrames - sched.longTermQueueSize)
  // log.Printf("[Long] Using queue of size %d\n", sched.longTermQueueSize)
  // log.Printf("[Long] A cushion of %d will be left in RAM\n", leftoverSpace)
  sched.EachWhile(func (p *process.Process) bool {
    avail := sched.pm.AvailableRAM()
    if avail < leftoverSpace {
      // the threshhold has been met (need to use the RAM for other stuff)
      // we should break out now and leave things as they are
      // log.Printf("[Long] Stopping at RAM %d/%d\n", avail, ivm.RAMNumFrames)
      return false
    }
    // fill the RAM with another process's initial footprint
    if err := sched.pm.Load(p); err != nil {
      msg := fmt.Sprintf("failed to load process %v", *p)
      log.Printf(
        "[Long] RAM is at %d/%d",
        sched.pm.AvailableRAM(),
        ivm.RAMNumFrames,
      )
      log.Printf("[Long] %s\n", msg)
      panic(msg)
      return false
    }
    log.Printf(
      "[Long] process %d is loaded (RAM is now %d/%d)\n",
      p.ProcessNumber, sched.pm.AvailableRAM(), ivm.RAMNumFrames,
    )
    return true
  })
}

package scheduler

import "../process"

// Short performs short-term scheduling operations.
func (sched Scheduler) Short(corenum uint8, p *process.Process) {
  // make sure the given core state is ready for business
  p.Status = process.Run
  p.CPUID = corenum

  // fill the caches with the appropriate contents of RAM
  p.State.Caches = sched.pm.CachesForProcess(p)
}

package scheduler

import (
	"../../vm/ivm"
	"../process"
)

// Long performs long-term scheduling operations.
// This means adding any processes to RAM if possible.
func (sched *Scheduler) Long() {
	sched.logger.Printf("running long scheduler")

	// go through each process until the RAM is sufficiently filled
	// this is a tunable value, of course (since we want some spare space)
	leftoverSpace := int(ivm.RAMNumFrames - sched.longTermQueueSize)
	sched.logger.Printf(
		"using queue of size %d (a cushion of %d will be left)",
		sched.longTermQueueSize, leftoverSpace,
	)
	sched.EachWhile(func(p *process.Process) bool {
		avail := sched.pm.AvailableRAM()
		if avail < leftoverSpace {
			// the threshhold has been met (need to use the RAM for other stuff)
			// we should break out now and leave things as they are
			sched.logger.Printf(
				"stopping long-term scheduling (RAM at %d/%d)",
				avail, ivm.RAMNumFrames,
			)
			return false
		}
		// fill the RAM with another process's initial footprint
		if err := (*sched.pm).Load(p); err != nil {
			sched.logger.Printf(
				"long-term scheduling must stop (RAM is at %d/%d)",
				sched.pm.AvailableRAM(), ivm.RAMNumFrames,
			)
			sched.logger.Panicf(
				"failed to load process %d during long-term scheduling: %v",
				p.ProcessNumber, err,
			)
			return false
		}
		sched.logger.Printf(
			"process %d has been loaded for long-term scheduling (RAM is at %d/%d)",
			p.ProcessNumber, sched.pm.AvailableRAM(), ivm.RAMNumFrames,
		)
		return true
	})
}

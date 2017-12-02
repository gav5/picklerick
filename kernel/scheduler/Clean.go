package scheduler

import (
	"../../vm/ivm"
	"../process"
)

// Clean ensures spaces in RAM that are no longer used are given back.
// (this, of course, means terminated processes)
func (sched *Scheduler) Clean() error {
	sched.logger.Printf(
		"[Clean] %d items in the process unload queue (RAM at %d/%d)",
		sched.processUnloadQueue.Len(), sched.pm.AvailableRAM(), ivm.RAMNumFrames,
	)

	// unload each process in the unload queue
	for sched.processUnloadQueue.Len() > 0 {
		pNum := sched.processUnloadQueue.Peek()
		p := sched.Find(pNum)
		var err error

		if p == nil {
			sched.logger.Panicf(
				"[Clean] process %d in the process list (it is in the unload queue)",
				pNum,
			)
		} else if p.Status() != process.Terminated {
			sched.logger.Panicf(
				"[Clean] process %d not marked as terminated (it is in the unload queue)",
				pNum,
			)
		} else {
			sched.logger.Printf("[Clean] going to clean process %d", pNum)
		}

		// make sure to save the process (so progress isn't lost)
		err = sched.Save(p)
		if err != nil {
			// put this back onto the process unload queue (to be handled later)
			// this is okay to do since we are returning here anyway
			// sched.processUnloadQueue.Push(pNum)

			sched.logger.Printf(
				"[Clean] ERROR saving process %d: %v",
				pNum, err,
			)
			return err
		}

		// only after we have successfully saved can we unload the data from RAM
		// (otherwise we'd lose the results of the program)
		err = sched.Unload(p)
		if err != nil {
			// put this back onto the process unload queue (to be handled later)
			// this is okay to do since we are returning here anyway
			// sched.processUnloadQueue.Push(pNum)

			sched.logger.Printf(
				"[Clean] ERROR unloading process %d: %v",
				pNum, err,
			)
			return err
		}

		_ = sched.processUnloadQueue.Pop()
		sched.logger.Printf(
			"[Clean] process %d has been cleaned",
			pNum,
		)
	}
	return nil
}

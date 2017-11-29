package scheduler

import "../process"

// Clean ensures spaces in RAM that are no longer used are given back.
// (this, of course, means terminated processes)
func (sched *Scheduler) Clean() error {
	sched.logger.Printf("clean RAM")

	// unload each process in the unload queue
	for sched.processUnloadQueue.Len() > 0 {
		pNum := sched.processUnloadQueue.Pop().(uint8)
		p := sched.Find(pNum)
		var err error

		if p == nil {
			sched.logger.Panicf(
				"process %d not found in process list (it is in the unload queue)",
				pNum,
			)
		} else if p.Status() != process.Terminated {
			sched.logger.Panicf(
				"process %d is not marked as terminated (it is in the unload queue)",
				pNum,
			)
		} else {
			sched.logger.Printf("going to clean process %d", pNum)
		}

		// make sure to save the process (so progress isn't lost)
		err = sched.Save(p)
		if err != nil {
			// put this back onto the process unload queue (to be handled later)
			// this is okay to do since we are returning here anyway
			sched.processUnloadQueue.Push(pNum)

			sched.logger.Printf(
				"ERROR saving process %d during clean procedure: %v",
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
			sched.processUnloadQueue.Push(pNum)

			sched.logger.Printf(
				"ERROR unloading process %d during clean procedure: %v",
				pNum, err,
			)
			return err
		}

		sched.logger.Printf(
			"process %d has been saved and unloaded successfully!",
			pNum,
		)
	}
	return nil
}

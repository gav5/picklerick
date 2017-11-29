package scheduler

import "../process"

// Clean ensures spaces in RAM that are no longer used are given back.
// (this, of course, means terminated processes)
func (sched *Scheduler) Clean() error {
  sched.logger.Printf("clean RAM")

  // go through each process and handle the terminated ones
  return sched.EachWithError(func(p *process.Process) error {
    if p.Status() == process.Terminated {
      sched.logger.Printf(
        "process %d TERMINATED => should save and unload",
        p.ProcessNumber,
      )

      var err error
      // this process is terminated! (so it should't have any RAM)
      // make sure to save the process (so progress isn't lost)
      err = sched.Save(p)
      if err != nil {
        sched.logger.Printf(
          "ERROR saving process %d during clean procedure: %v",
          p.ProcessNumber, err,
        )
        return err
      }
      // only after we have successfully saved can we unload the data from RAM
      // (otherwise we'd lose the results of the program)
      err = sched.Unload(p)
      if err != nil {
        sched.logger.Printf(
          "ERROR unloading process %d during clean procedure: %v",
          p.ProcessNumber, err,
        )
        return err
      }
      sched.logger.Printf(
        "process %d has been saved and unloaded successfully!",
        p.ProcessNumber,
      )
    }
    return nil
  })
}

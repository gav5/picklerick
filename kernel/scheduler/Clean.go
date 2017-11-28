package scheduler

import "../process"

// Clean ensures spaces in RAM that are no longer used are given back.
// (this, of course, means terminated processes)
func (sched *Scheduler) Clean() error {
  // go through each process and handle the terminated ones
  return sched.EachWithError(func(p *process.Process) error {
    if p.Status() == process.Terminated {
      var err error
      // this process is terminated! (so it should't have any RAM)
      // make sure to save the process (so progress isn't lost)
      err = sched.Save(p)
      if err != nil {
        return err
      }
      // only after we have successfully saved can we unload the data from RAM
      // (otherwise we'd lose the results of the program)
      err = sched.Unload(p)
      if err != nil {
        return err
      }
    }
    return nil
  })
}

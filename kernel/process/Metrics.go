package process

import (
  "../../metric/stopwatch"
  "../../metric/fractional"
  "../../metric/counter"
)

// Metrics is used to collect information about the process.
type Metrics struct {
  JobWaitTime       stopwatch.Stopwatch
  JobCompletionTime stopwatch.Stopwatch
  IOOperationCount  counter.Counter
  RAMUse            fractional.Fractional
  CacheUse          fractional.Fractional
}

type statusTransition struct {
  from Status
  to Status
}

// ReactToStatus reacts to a change in state.
func (m *Metrics) ReactToStatus(o, n Status) {
  tr := statusTransition{o, n}
  switch tr {
  case statusTransition{New, Ready}:
    // we are now waiting and tracking execution
    m.JobCompletionTime.Start()
    m.JobWaitTime.Start()
  case statusTransition{Ready, Run}:
    // we are not waiting anymore
    m.JobWaitTime.Stop()
  case statusTransition{Run, Wait}:
    // we are now waiting again
    m.JobWaitTime.Start()
  case statusTransition{Run, Terminated}:
    // we have completed the job!
    m.JobCompletionTime.Stop()
  }
}

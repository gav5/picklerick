package process

// Status indicates what a given process is doing at the moment.
type Status uint8

const (
  // New means the process has just been created.
  New Status = iota
  // Ready means the process is in the queue and ready to be run.
  Ready
  // Run means the process is currently being run.
  Run
  // Wait means the process is waiting on memory to be filled.
  Wait
  // Terminated means the process has completed its course.
  Terminated
)

func (status Status) String() string {
  switch status {
  case New:
    return "New"
  case Ready:
    return "Ready"
  case Run:
    return "Run"
  case Wait:
    return "Wait"
  case Terminated:
    return "Terminated"
  default:
    return "???"
  }
}

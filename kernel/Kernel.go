package kernel

import (
  "../vm/ivm"
  "../config"
  "../util/logger"
  "./pageManager"
  "./scheduler"
  "./process"
  "./loader"
  "log"
  "io"
)

// Kernel houses all the storage and functionality of the OS kernel.
type Kernel struct {
  config config.Config
  virtualMachine ivm.IVM
  pm pageManager.PageManager
  sched *scheduler.Scheduler
  logger *log.Logger
}

// New makes a kernel with the given virtual machine.
func New(virtualMachine ivm.IVM, c config.Config) (*Kernel, error) {

  k := &Kernel{
    config: c,
    virtualMachine: virtualMachine,
    pm: pageManager.Make(virtualMachine),
    logger: logger.New("kernel"),
  }

  // load the programs from file (via loader)
  programs, err := loader.Load(c.Progfile)
  if err != nil {
    return nil, err
  }
  k.logger.Printf("Got %d programs!\n", len(programs))
  k.sched = scheduler.New(c, &k.pm, programs)

  return k, nil
}

// Tick is used to signal the start of a virtual machine cycle to the kernel.
// This sets up processes and resources before the next cycle begins.
func (k Kernel) Tick() {
  // defer to the scheduler
  k.logger.Println("Tick!")
  k.sched.Tick()
}

// Tock is used to signal the end of a virtual machine cycle to the kernel.
// This reacts to the events that occured during the cycle.
func (k Kernel) Tock() error {
  // defer to the Scheduler
  k.logger.Println("Tock!")
  err := k.sched.Tock()
  if err != nil {
    k.logger.Printf("Tock Error: %v", err)
  }
  return err
}

// ProcessForCore returns the appropriate process for the given core.
func (k Kernel) ProcessForCore(corenum uint8) process.Process {
  // defer to the scheduler
  p := k.sched.ProcessForCore(corenum)
  k.logger.Printf(
    "giving process %d to core %d",
    p.ProcessNumber, corenum,
  )
  return p
}

// UpdateProcess updates an existing process in the list.
func (k Kernel) UpdateProcess(p process.Process) error {
  // defer to the scheduler
  k.logger.Printf(
    "updating process %d (status %v)",
    p.ProcessNumber, p.Status,
  )
  err := k.sched.Update(p)
  if err != nil {
    k.logger.Printf(
      "ERROR updating process %d: %v",
      p.ProcessNumber, err,
    )
  }
  return err
}

// LoadProcess makes sure the given process is in RAM.
func (k Kernel) LoadProcess(p *process.Process) error {
  // defer to the scheduler
  k.logger.Printf("loading process %d", p.ProcessNumber)
  err := k.sched.Load(p)
  if err != nil {
    k.logger.Printf(
      "ERROR loading process %d: %v",
      p.ProcessNumber, err,
    )
  }
  return err
}

// CompleteProcess marks a process as completed and removes its used resources.
// (this gives the system the opportunity to fill those resources for others)
func (k Kernel) CompleteProcess(p *process.Process) {
  // defer to the scheduler
  k.logger.Printf("completing process %d", p.ProcessNumber)
  k.sched.Complete(p)
}

// IsDone returns if the system is done yet.
func (k Kernel) IsDone() bool {
  return k.sched.IsDone()
}

// NumProcessesLeft returns the number of processes still left in the queue.
func (k Kernel) NumProcessesLeft() int {
  return k.sched.NumLeft()
}

// FprintProcessTable prints the process table to the given writer.
func (k Kernel) FprintProcessTable(w io.Writer) error {
  return k.sched.FprintProcessTable(w)
}

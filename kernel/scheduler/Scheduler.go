package scheduler

import (
  "container/heap"
  "sort"
  "io"
  "fmt"
  // "log"

  "../../config"
  "../process"
  "../program"
  "../pageManager"
)

// Scheduler keeps track of system processes.
type Scheduler struct {
  processList *processList
  completed []process.Process
  pm *pageManager.PageManager
  methodName string
  longTermQueueSize uint
}

// New creates a new scheduler.
func New(c config.Config, p *pageManager.PageManager, a []program.Program) *Scheduler {
  sched := &Scheduler{
    processList: &processList{
      base: process.MakeArray(a),
      sortMethod: MethodForSwitch(c.Sched),
    },
    completed: []process.Process{},
    pm: p,
    methodName: c.Sched,
    longTermQueueSize: c.QSize,
  }

  // sort the whole thing
  sort.Sort(sched.processList)

  // make sure each process is set up with the page manager
  sched.Each(func(p *process.Process) {
    sched.pm.Setup(p)
  })
  return sched
}

// Tick is used to signal the start of a virtual machine cycle to the kernel.
// This sets up processes and resources before the next cycle begins.
func (sched Scheduler) Tick() {
  // TODO: handle processes that shouldn't be on a CPU anymore
  // TODO: handle processes that shouldn't be in RAM anymore
  // run the long-term scheduler
  sched.Long()
}

// Tock is used to signal the end of a virtual machine cycle to the kernel.
// This reacts to the events that occured during the cycle.
func (sched Scheduler) Tock() {
  // check back through previous requests and try to fulfill them
  sched.pm.HandleWaitlist()
  // make sure any waiting processes have what they need
  sched.Each(func(p *process.Process) {
    if p.Status == process.Wait {
      err := sched.pm.Reallocate(p)
      if err == nil {
        p.Status = process.Ready
      }
    }
  })
}

// ProcessForCore returns the appropriate process for the given core.
func (sched Scheduler) ProcessForCore(corenum uint8) *process.Process {
  // Look for the first process that is ready to be run
  p := sched.FindBy(func(p *process.Process) bool {
    return p.Status == process.Ready
  })
  if p != nil {
    // make sure the process is ready to be run on the given core
    sched.Short(corenum, p)
  }
  return p
}

// Update updates an existing process in the list.
func (sched *Scheduler) Update(p *process.Process) error {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    pX := &sched.processList.base[i]
    if p == pX {
      heap.Fix(sched.processList, i)
      return nil
    }
  }
  return NotFoundError{}
}

// Load makes sure the given process is in RAM.
func (sched *Scheduler) Load(p *process.Process) error {
  if p.Status != process.Ready {
    // this process is not ready to be loaded into RAM
    return NotReadyError{}
  }
  // defer to the page manager
  sched.pm.Load(p)
  return nil
}

// Complete completes the given processs (marks it Terminated).
func (sched *Scheduler) Complete(p *process.Process) error {
  index, _ := sched.findPair(func(px *process.Process) bool {
    return p == px
  })
  if index == -1 {
    return NotFoundError{}
  }
  // remove from the process list
  sched.processList.base = append(
    sched.processList.base[:index],
    sched.processList.base[index+1:]...,
  )
  // add to the completed list
  sched.completed = append(sched.completed, *p)
  return nil
}

// NotFoundError is when the desired process is not in the list.
type NotFoundError struct{}
func (err NotFoundError) Error() string {
  return "process is not in the scheduler"
}

// NotReadyError is when the process is not ready to load RAM.
type NotReadyError struct{}
func (err NotReadyError) Error() string {
  return "process is not ready to load RAM"
}

// Each goes through each process in order and passes to the given function.
func (sched Scheduler) Each(fn func(*process.Process)) {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    fn(&sched.processList.base[i])
  }
}

// EachWithError goes through each process and checks for an error each time.
func (sched Scheduler) EachWithError(fn func(*process.Process) error) error {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    if err := fn(&sched.processList.base[i]); err != nil {
      return err
    }
  }
  return nil
}

// EachWhile goes through each process while the function keeps returning true.
func (sched Scheduler) EachWhile(fn func(*process.Process) bool) {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    if !fn(&sched.processList.base[i]) {
      break
    }
  }
}

// FindBy goes through each until the passed function returns true.
func (sched Scheduler) FindBy(fn func(*process.Process) bool) *process.Process {
  _, p := sched.findPair(fn)
  return p
}

func (sched Scheduler) findPair(fn func(*process.Process) bool) (int, *process.Process) {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    p := &sched.processList.base[i]
    if fn(p) {
      return i, p
    }
  }
  return -1, nil
}

// FprintProcessTable prints the process table to the given writer.
func (sched Scheduler) FprintProcessTable(w io.Writer) error {
  combined := append(sched.processList.base, sched.completed...)
  header := fmt.Sprintf(
    "Process Table (%d processes, sort method: %s, queue size: %d)\n",
    len(combined), sched.methodName, sched.longTermQueueSize,
  )
  if _, err := w.Write([]byte(header)); err != nil {
    return err
  }
  for i := len(combined)-1; i >= 0; i-- {
    p := combined[i]
    out := fmt.Sprintf(
      "[%02d] %-10s p%02d (%d instructions) {RAM: %2d pages} {Disk: %2d pages}\n",
      p.ProcessNumber, p.Status, p.Priority,
      p.CodeSize, len(p.RAMPageTable), len(p.DiskPageTable),
    )
    if _, err := w.Write([]byte(out)); err != nil {
      return err
    }
  }
  return nil
}

// IsDone returns if the system is done yet.
func (sched Scheduler) IsDone() bool {
  return sched.processList.Len() == 0
}

// NumLeft is the number of processes left in the queue.
func (sched Scheduler) NumLeft() int {
  return sched.processList.Len()
}

// Add adds a process into the process manager.
func (sched Scheduler) Add(p process.Process) {
  heap.Push(sched.processList, p)
}

// Find returns the process with the corresponding process number.
func (sched Scheduler) Find(processNumber uint8) *process.Process {
  _, p := sched.findPair(func(p *process.Process) bool {
    return p.ProcessNumber == processNumber
  })
  return p
}

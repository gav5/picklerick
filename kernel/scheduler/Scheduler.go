package scheduler

import (
  "container/heap"
  "sort"
  "io"
  "fmt"
  "strings"
  "log"

  "../../config"
  "../process"
  "../program"
  "../pageManager"
  "../../vm/ivm"
)

// Scheduler keeps track of system processes.
type Scheduler struct {
  processList *processList
  completed *processList
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
    completed: &processList{
      base: []process.Process{},
      sortMethod: FCFS,
    },
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
  // run the long-term scheduler
  sched.Long()
}

// Tock is used to signal the end of a virtual machine cycle to the kernel.
// This reacts to the events that occured during the cycle.
func (sched Scheduler) Tock() error {

  // make sure terminated processes aren't taking up space anymore
  // (otherwise, there's nothing to fill here and it just stops)
  err := sched.Clean()
  if err != nil {
    return err
  }

  // check back through previous requests and try to fulfill them
  sched.pm.HandleWaitlist()

  // make sure any waiting processes have what they need
  sched.Each(func(p *process.Process) {
    if p.Status() == process.Wait {
      err := sched.pm.Reallocate(p)
      if err == nil {
        p.SetStatus(process.Ready)
      }
    }
  })
  return nil
}

// ProcessForCore returns the appropriate process for the given core.
func (sched Scheduler) ProcessForCore(corenum uint8) process.Process {
  // Look for the first process that is ready to be run
  p := sched.FindBy(func(p *process.Process) bool {
    return p.Status() == process.Ready
  })
  if p == nil {
    return process.Sleep()
  }
  // make sure the process is ready to be run on the given core
  sched.Short(corenum, p)
  // update this internally (because Short changed it)
  sched.Update(*p)

  return *p
}

// Update updates an existing process in the list.
func (sched *Scheduler) Update(p process.Process) error {
  for i := sched.processList.Len() - 1; i >= 0; i-- {
    pX := sched.processList.base[i]
    if p.ProcessNumber == pX.ProcessNumber {
      sched.processList.base[i] = p
      return nil
    }
  }
  return NotFoundError{}
}

// Load makes sure the given process is in RAM.
func (sched *Scheduler) Load(p *process.Process) error {
  if p.Status() != process.Ready {
    // this process is not ready to be loaded into RAM
    return NotReadyError{}
  }
  // defer to the page manager
  err := sched.pm.Load(p)
  if err != nil {
    return err
  }
  return nil
}

// Save makes sure the given process's RAM is persisted to Disk.
func (sched *Scheduler) Save(p *process.Process) error {
  // defer to the page manager
  return sched.pm.Save(p)
}

// Unload makes sure the given process is not in RAM.
func (sched *Scheduler) Unload(p *process.Process) error {
  log.Printf("[Unload] process %d should be unloaded\n", p.ProcessNumber)

  if p.Status() != process.Terminated {
    log.Panicf(
      "[Unload] process %d is not terminated (is %v)",
      p.ProcessNumber, p.Status,
    )
  }

  // defer to the page manager
  err := sched.pm.Unload(p)
  if err != nil {
    return err
  }

  // make sure the process is now removed (for efficiency)
  // we will want to add to the completed list
  return sched.removeProcess(p)
}

func (sched *Scheduler) removeProcess(p *process.Process) error {
  index, _ := sched.findPair(func(px *process.Process) bool {
    return p.ProcessNumber == px.ProcessNumber
  })
  if index == -1 {
    return NotFoundError{}
  }
  log.Printf(
    "[removeProcess] (before: %d) %d completed\n",
    p.ProcessNumber, sched.completed.Len(),
  )

  // remove from the process list
  (*sched.processList).base = append(
    sched.processList.base[:index],
    sched.processList.base[index+1:]...,
  )
  // add to the completed list
  sched.completed.Push(*p)
  sort.Sort(sched.completed)
  log.Printf(
    "[removeProcess] (after: %d) %d completed\n",
    p.ProcessNumber, sched.completed.Len(),
  )
  return nil
}

// Complete completes the given processs (marks it Terminated).
func (sched *Scheduler) Complete(p *process.Process) error {
  // mark is Terminated (this will get cleaned up later)
  p.SetStatus(process.Terminated)
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
  todoLen := sched.processList.Len()
  completedLen := sched.completed.Len()
  combinedLen := completedLen + todoLen
  header := fmt.Sprintf(
    "Process Table (%d/%d processes, sort method: %s, queue size: %d/%d)\n",
    todoLen, combinedLen, sched.methodName,
    sched.longTermQueueSize, ivm.RAMNumFrames,
  )
  if _, err := w.Write([]byte(header)); err != nil {
    return err
  }
  if err := sched.completed.fprint(w); err != nil {
    return err
  }
  bars := fmt.Sprintf("%s\n", strings.Repeat("=", 70))
  if _, err := w.Write([]byte(bars)); err != nil {
    return err
  }
  if err := sched.processList.fprint(w); err != nil {
    return err
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

package kernel

import (
  "../vm/ivm"
  "../config"
  "./pageManager"
  "./scheduler"
  "./process"
  // "./page"
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
}

// New makes a kernel with the given virtual machine.
func New(virtualMachine ivm.IVM, c config.Config) (*Kernel, error) {
  // load the programs from file (via loader)
  programs, err := loader.Load(c.Progfile)
  if err != nil {
    return nil, err
  }
  log.Printf("Got %d programs!\n", len(programs))

  k := &Kernel{
    config: c,
    virtualMachine: virtualMachine,
    pm: pageManager.Make(virtualMachine),
  }
  k.sched = scheduler.New(c.Sched, &k.pm, programs)

  return k, nil
}

// Tick is used to signal the start of a virtual machine cycle to the kernel.
// This sets up processes and resources before the next cycle begins.
func (k Kernel) Tick() {
  // defer to the scheduler
  k.sched.Tick()
}

// Tock is used to signal the end of a virtual machine cycle to the kernel.
// This reacts to the events that occured during the cycle.
func (k Kernel) Tock() {
  // TODO: make Tock()
}

// ProcessForCore returns the appropriate process for the given core.
func (k Kernel) ProcessForCore(corenum uint8) *process.Process {
  // defer to the scheduler
  return k.sched.ProcessForCore(corenum)
}

// UpdateProcess updates an existing process in the list.
func (k Kernel) UpdateProcess(p *process.Process) error {
  // defer to the scheduler
  return k.sched.Update(p)
}

// LoadProcess makes sure the given process is in RAM.
func (k Kernel) LoadProcess(p *process.Process) error {
  // defer to the scheduler
  return k.sched.Load(p)
}

// CompleteProcess marks a process as completed and removes its used resources.
// (this gives the system the opportunity to fill those resources for others)
func (k Kernel) CompleteProcess(p *process.Process) {
  // defer to the scheduler
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

// PageReadRAM reads a page from the RAM at the given page number and page table.
// func (k Kernel) PageReadRAM(pageNumber page.Number, pageTable page.Table) page.Page {
// 	frameNumber := pageTable[pageNumber]
// 	frame := k.virtualMachine.RAMFrameFetch(frameNumber)
// 	return page.Page(frame)
// }
//
// // PageWriteRAM writes the given page to the RAM at the given page number.
// func (k Kernel) PageWriteRAM(p page.Page, pageNumber page.Number, pageTable page.Table) {
// 	frameNumber := pageTable[pageNumber]
// 	frame := ivm.Frame(p)
// 	k.virtualMachine.RAMFrameWrite(frameNumber, frame)
// }
//
// // PageReadDisk reads a page from the Disk at the given page number and page table.
// func (k Kernel) PageReadDisk(pageNumber page.Number, pageTable page.Table) page.Page {
//   frameNumber := pageTable[pageNumber]
//   frame := k.virtualMachine.DiskFrameFetch(frameNumber)
//   return page.Page(frame)
// }
//
// // PageWriteDisk writes the given page to the Disk at the given page numbber.
// func (k Kernel) PageWriteDisk(p page.Page, pageNumber page.Number, pageTable page.Table) {
//   frameNumber := pageTable[pageNumber]
//   frame := ivm.Frame(p)
//   k.virtualMachine.DiskFrameWrite(frameNumber, frame)
// }

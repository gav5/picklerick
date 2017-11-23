package kernel

import (
  "../vm/ivm"
  "../prog"
  "../config"
  "./resourceManager"
  "./processManager"
  "./process"
  "./page"
  "./scheduler"
  "log"
)

// Kernel houses all the storage and functionality of the OS kernel.
type Kernel struct {
  config config.Config
  virtualMachine ivm.IVM
  ramRM *resourceManager.ResourceManager
  diskRM *resourceManager.ResourceManager
  pm *processManager.ProcessManager
}

// New makes a kernel with the given virtual machine.
func New(virtualMachine ivm.IVM, c config.Config) (*Kernel, error) {
  k := &Kernel{
    config: c,
    virtualMachine: virtualMachine,
    ramRM: resourceManager.New(ivm.RAMNumFrames),
    diskRM: resourceManager.New(ivm.DiskNumFrames),
    pm: processManager.New(scheduler.FIFO),
  }
  // load programs into the system
  var programArray []prog.Program
  var err error
	if programArray, err = prog.ParseFile(c.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return k, err
	}
  log.Printf("Got %d programs!\n", len(programArray))
  if err = k.LoadPrograms(programArray); err != nil {
    return k, err
  }
  return k, nil
}

// ProcessForCore returns the appropriate process for the given core.
// func (k Kernel) ProcessForCore(coreNum int) *process.Process {
//   return k.pm.ProcessForCore(coreNum)
// }

// Tock should be called after every cycle completes
// func (k Kernel) Tock() {
//   k.pm.Reevaluate()
// }

// PopProcess removes the frontmost process off the queue and returns it.
func (k Kernel) PopProcess() process.Process {
  return k.pm.Pop()
}

// IsDone returns if the system is done yet.
func (k Kernel) IsDone() bool {
  return k.pm.IsDone()
}

// NumProcessesLeft returns the number of processes still left in the queue.
func (k Kernel) NumProcessesLeft() int {
  return k.pm.NumLeft()
}

// PageReadRAM reads a page from the RAM at the given page number and page table.
func (k Kernel) PageReadRAM(pageNumber page.Number, pageTable page.Table) page.Page {
	frameNumber := pageTable[pageNumber]
	frame := k.virtualMachine.RAMFrameFetch(frameNumber)
	return page.Page(frame)
}

// PageWriteRAM writes the given page to the RAM at the given page number.
func (k Kernel) PageWriteRAM(p page.Page, pageNumber page.Number, pageTable page.Table) {
	frameNumber := pageTable[pageNumber]
	frame := ivm.Frame(p)
	k.virtualMachine.RAMFrameWrite(frameNumber, frame)
}

// PageReadDisk reads a page from the Disk at the given page number and page table.
func (k Kernel) PageReadDisk(pageNumber page.Number, pageTable page.Table) page.Page {
  frameNumber := pageTable[pageNumber]
  frame := k.virtualMachine.DiskFrameFetch(frameNumber)
  return page.Page(frame)
}

// PageWriteDisk writes the given page to the Disk at the given page numbber.
func (k Kernel) PageWriteDisk(p page.Page, pageNumber page.Number, pageTable page.Table) {
  frameNumber := pageTable[pageNumber]
  frame := ivm.Frame(p)
  k.virtualMachine.DiskFrameWrite(frameNumber, frame)
}

// PushProgram pushes a program into the first available space in the VM.
// (this prefers RAM, but falls back to using the disk if necessary; returns error otherwise)
func (k *Kernel) PushProgram(p prog.Program, pageTable *page.Table) error {
	// get the pages for the given program
	// push those pages into the VM and return the result
  pageArray := page.ArrayFromProgram(p)
  numPages := len(pageArray)
  if numPages <= k.ramRM.QuantityAvailable() {
    // there is adequate space in RAM to fit this!!
    // let's claim some space in RAM and put it in there!
    frameNumbers, err := k.ramRM.Claim(numPages)
    if err != nil {
      return err
    }
    // add to the page table (so the process knows to use that space)
    for pnum, fnum := range frameNumbers {
      (*pageTable)[page.Number(pnum)] = ivm.FrameNumber(fnum)
    }
    // insert the given pages into RAM
    for pageNumber, p := range pageArray {
      k.PageWriteRAM(p, page.Number(pageNumber), *pageTable)
    }
  }
  return nil
}

// PushOverflowError means there isn't enough storage to hold all provided data.
type PushOverflowError struct{}

func (e PushOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided data."
}

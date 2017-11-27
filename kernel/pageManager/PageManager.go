package pageManager

import (
  "log"
  "fmt"

  "../../vm/ivm"
  "../page"
  "../process"
  "../resourceManager"
)

// PageManager is responsible for assigning pages of memory out to RAM.
type PageManager struct {
  virtualMachine ivm.IVM
  ramRM *resourceManager.ResourceManager
  diskRM *resourceManager.ResourceManager
  waitlist []*process.Process
}

// Make builds a new PageManager instance.
func Make(virtualMachine ivm.IVM) PageManager {
  pm := PageManager{
    virtualMachine: virtualMachine,
    ramRM: resourceManager.New(ivm.RAMNumFrames),
    diskRM: resourceManager.New(ivm.DiskNumFrames),
    waitlist: []*process.Process{},
  }
  return pm
}

// Setup creates new space on disk for the given process.
func (pm *PageManager) Setup(p *process.Process) error {
  // determine the initial pages needed here
  instructions := page.ArrayFromUint32Array(p.Program.Instructions)
  data := page.ArrayFromUint32Array(p.Program.DataBlock)[:p.Footprint]
  initialPages := append(instructions, data...)

  // claim frames for the initial pages used
  frameNumbers, err := pm.diskRM.Claim(len(initialPages))
  if err != nil {
    return err
  }

  // set the (disk) page table for the process
  // write appropriate page content to those frames
  for i, x := range frameNumbers {
    pn := page.Number(i)
    fn := ivm.FrameNumber(x)
    p.DiskPageTable[pn] = fn
    frame := ivm.Frame(initialPages[i])
    pm.virtualMachine.DiskFrameWrite(fn, frame)
  }
  return nil
}

// AvailableRAM returns the number of available frames in RAM
func (pm PageManager) AvailableRAM() int {
  return pm.ramRM.QuantityAvailable()
}

// CachesForProcess returns the appropriate caches for the given process.
func (pm PageManager) CachesForProcess(p *process.Process) ivm.FrameCache {
  caches := make(ivm.FrameCache)
  for pn, fn := range p.RAMPageTable {
    caches[ivm.FrameNumber(pn)] = pm.virtualMachine.RAMFrameFetch(fn)
  }
  return caches
}

// Load makes sure the given pages are in RAM.
func (pm *PageManager) Load(p *process.Process) error {
  // the initial claim will be the instructions and data footprint
  // initialClaim := len(p.Program.Instructions) + p.Footprint
  initialClaim := len(p.DiskPageTable)
  frameNums, err := pm.ramRM.Claim(initialClaim)
  if err != nil {
    return err
  }
  // assign to the RAM page table and copy over to designated frames
  for i, x := range frameNums {
    // assign to page table
    pn := page.Number(i)
    rfn := ivm.FrameNumber(x)
    p.RAMPageTable[pn] = rfn
    // copy frame from disk to RAM
    dfn := p.DiskPageTable[pn]
    frame := pm.virtualMachine.DiskFrameFetch(dfn)
    pm.virtualMachine.RAMFrameWrite(rfn, frame)
  }
  return nil
}

// Save makes sure the given process's RAM is persisted to Disk.
func (pm *PageManager) Save(p *process.Process) error {
  // go through each page in RAM and persist to the corresponding page on Disk
  // if that page is not on Disk yet, we will have to make one first
  for pn, rfn := range p.RAMPageTable {
    if _, present := p.DiskPageTable[pn]; !present {
      // there is no corresponding page yet, so make sure there is one
      // these will be claimed one at a time (it can get back to it later)
      newDfn, err := pm.diskRM.Claim(1)
      if err != nil {
        return err
      }
      // assign the new disk frame number to the disk page table
      p.DiskPageTable[pn] = ivm.FrameNumber(newDfn[0])
    }
    // write the frame from RAM to the corresponding page on Disk
    // (note it should be there now becasue of the above guard)
    ramFrame := pm.virtualMachine.RAMFrameFetch(rfn)
    dfn, pnOk := p.DiskPageTable[pn]
    if !pnOk {
      // just to be safe!
      panic(fmt.Sprintf(
        "tried to fetch a page number (%d) that wasn't there!", pn,
      ))
    }
    pm.virtualMachine.DiskFrameWrite(dfn, ramFrame)
  }
  return nil
}

// Unload makes sure the given process is not in RAM.
func (pm *PageManager) Unload(p *process.Process) error {
  // at some point, we're going to have to remove some page table entries
  // we're also going to need to release some frames from the resource manager
  ptLen := len(p.RAMPageTable)
	pgNumbers := make([]page.Number, ptLen)
	frNumbers := make([]int, ptLen)

  // go through each page in the RAM page table and zero them out
  // (so the next process's to get these frames have a clean slate)
  for pn, fn := range p.RAMPageTable {
    pm.virtualMachine.RAMFrameWrite(fn, ivm.MakeFrame())
    // while we're at it, let's fill the array
    pgNumbers[ptLen - 1] = pn
		frNumbers[ptLen - 1] = int(fn)
		ptLen--
  }

  // give back the frames to the RAM resource manager
  // (so it can go to some other process at some point)
  err := pm.ramRM.Release(frNumbers)
  if err != nil {
    return err
  }

  // remove the corresponing entries from the RAM page table
  // (this is done this way to ensure an entry wasn't missed)
  for _, pn := range pgNumbers {
    delete(p.RAMPageTable, pn)
  }

  // make sure we got all the entries!
  // if not, this should panic (becasue it's unexpected)
  if len(p.RAMPageTable) > 0 {
    log.Printf(
      "[Unload] %d page table entries still remain!?\n",
      len(p.RAMPageTable),
    )
    panic("RAM page table is not completely cleared!")
  }

  return nil
}

// Reallocate ensures the given process has enough space allocated to it.
// (basically, it handles pages-faults)
func (pm *PageManager) Reallocate(p *process.Process) error {
  err := pm.reallocate(p)
  if err != nil {
    log.Printf(
      "[Reallocate] process %d reallocation error: %v\n",
      p.ProcessNumber, err,
    )
    // the request cannot be granted!
    // add the process to the waitlist
    pm.waitlist = append(pm.waitlist, p)
  } else {
    log.Printf(
      "[Reallocate] process %d reallocated: %v\n",
      p.ProcessNumber, p.RAMPageTable,
    )
  }
  return err
}

// HandleWaitlist ensures the items in the waitlist are handled eventually.
func (pm *PageManager) HandleWaitlist() {
  completed := []int{}
  for i, p := range pm.waitlist {
    err := pm.reallocate(p)
    if err == nil {
      log.Printf(
        "[HandleWaitlist] process %d reallocated!\n",
        p.ProcessNumber,
      )
      // remove from the waitlist (later)
      completed = append(completed, i)
    } else {
      log.Printf(
        "[HandleWaitlist] process %d error: %v\n",
        p.ProcessNumber, err,
      )
    }
  }
  // remove the completed processes from the waitlist
  // this must be done in reverse order to preserve the index integrity
  // (i.e. as you remove values, the indexes shift)
  for i := len(completed)-1; i >= 0; i-- {
    index := completed[i]
    // remove the indicated index from the waitlist
    pm.waitlist = append(pm.waitlist[:index], pm.waitlist[index+1:]...)
  }
}

func (pm *PageManager) reallocate(p *process.Process) error {
  numFaults := len(p.State.Faults)
  frameNums, err := pm.ramRM.Claim(numFaults)
  if err != nil {
    return err
  }
  p.Footprint += numFaults
  i := 0
  for x, v := range p.State.Faults {
    if !v {
      panic("expected true!")
    }
    pn := page.Number(x)
    fn := ivm.FrameNumber(frameNums[i])
    pm.virtualMachine.RAMFrameWrite(fn, ivm.MakeFrame())
    p.RAMPageTable[pn] = fn
    i++
  }
  p.State.Faults = ivm.FaultList{}
  return nil
}

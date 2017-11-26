package pageManager

import (
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

// Load makes sure the given pages are in RAM.
func (pm *PageManager) Load(p *process.Process) error {
  // the initial claim will be the instructions and data footprint
  initialClaim := len(p.Program.Instructions) + p.Footprint
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

// Reallocate ensures the given process has enough space allocated to it.
// (basically, it handles pages-faults)
func (pm *PageManager) Reallocate(p *process.Process) error {
  err := pm.reallocate(p)
  if err != nil {
    // the request cannot be granted!
    // add the process to the waitlist
    pm.waitlist = append(pm.waitlist, p)
  }
  return err
}

// HandleWaitlist ensures the items in the waitlist are handled eventually.
func (pm *PageManager) HandleWaitlist() {
  completed := []int{}
  for i, p := range pm.waitlist {
    err := pm.reallocate(p)
    if err == nil {
      // remove from the waitlist (later)
      completed = append(completed, i)
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
  frameNums, err := pm.diskRM.Claim(numFaults)
  if err != nil {
    return err
  }
  p.Footprint += numFaults
  for i, x := range p.State.Faults {
    pn := page.Number(x)
    fn := ivm.FrameNumber(frameNums[i])
    pm.virtualMachine.DiskFrameWrite(fn, ivm.MakeFrame())
    p.RAMPageTable[pn] = fn
  }
  return nil
}

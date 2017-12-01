package pageManager

import (
	"fmt"
	"log"

	"../../util/logger"
	"../../vm/ivm"
	"../page"
	"../process"
	"../resourceManager"
)

// PageManager is responsible for assigning pages of memory out to RAM.
type PageManager struct {
	virtualMachine ivm.IVM
	ramRM          *resourceManager.ResourceManager
	diskRM         *resourceManager.ResourceManager
	waitlist       []*process.Process
	logger         *log.Logger
}

// Make builds a new PageManager instance.
func Make(virtualMachine ivm.IVM) PageManager {
	pm := PageManager{
		virtualMachine: virtualMachine,
		ramRM:          resourceManager.New(ivm.RAMNumFrames),
		diskRM:         resourceManager.New(ivm.DiskNumFrames),
		waitlist:       []*process.Process{},
		logger:         logger.New("pageManager"),
	}
	return pm
}

// Setup creates new space on disk for the given process.
func (pm *PageManager) Setup(p *process.Process) error {
	// determine the initial pages needed here
	initialData := p.Program.RAMRepresentation()
	initialPages := page.ArrayFromUint32Array(initialData)

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
	pm.logger.Printf(
		"returning caches for process %d => %v",
		p.ProcessNumber, caches,
	)
	return caches
}

// Load makes sure the given pages are in RAM.
func (pm *PageManager) Load(p *process.Process) error {
	pm.logger.Printf("load process %d", p.ProcessNumber)
	initialClaim := len(p.DiskPageTable)
	pm.logger.Printf(
		"making initial claim of %d for process %d",
		initialClaim, p.ProcessNumber,
	)
	frameNums, err := pm.ramRM.Claim(initialClaim)
	if err != nil {
		pm.logger.Printf(
			"ERROR loading process %d: %v",
			p.ProcessNumber, err,
		)
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
	pm.logger.Printf(
		"RAM page table modified for process %d: %v",
		p.ProcessNumber, p.RAMPageTable,
	)
	return nil
}

// Apply makes sure the caches in the process are persisted to RAM.
func (pm *PageManager) Apply(p process.Process) error {
	pm.logger.Printf("[Apply] apply process %d", p.ProcessNumber)
	for pn, frame := range p.State().Caches {
		fn, ok := p.RAMPageTable.SoftFrameNumberForPageNumber(page.Number(pn))
		if !ok {
			// this should have been allocated already
			return fmt.Errorf(
				"page number %v does not exist in the page table for process %d",
				pn, p.ProcessNumber,
			)
		}
		pm.logger.Printf(
			"[Apply] writing to frame %v for process %d[%v]",
			fn, p.ProcessNumber, pn,
		)
		// write the given frame to RAM
		pm.virtualMachine.RAMFrameWrite(fn, frame)
	}
	return nil
}

// Save makes sure the given process's RAM is persisted to Disk.
func (pm *PageManager) Save(p *process.Process) error {
	pm.logger.Printf("save process %d", p.ProcessNumber)

	// go through each page in RAM and persist to the corresponding page on Disk
	// if that page is not on Disk yet, we will have to make one first
	for pn, rfn := range p.RAMPageTable {
		if _, present := p.DiskPageTable[pn]; !present {
			// there is no corresponding page yet, so make sure there is one
			// these will be claimed one at a time (it can get back to it later)
			newDfn, err := pm.diskRM.Claim(1)
			if err != nil {
				pm.logger.Printf(
					"ERROR making claim to disk for process %d: %v",
					p.ProcessNumber, err,
				)
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
			pm.logger.Panicf(
				"tried to fetch a page number (%d) that wasn't there!", pn,
			)
		}
		pm.virtualMachine.DiskFrameWrite(dfn, ramFrame)
	}
	pm.logger.Printf(
		"process %d saved to disk; page table now %v",
		p.ProcessNumber, p.DiskPageTable,
	)
	return nil
}

// Unload makes sure the given process is not in RAM.
func (pm *PageManager) Unload(p *process.Process) error {
	pm.logger.Printf("unload process %d", p.ProcessNumber)

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
		pgNumbers[ptLen-1] = pn
		frNumbers[ptLen-1] = int(fn)
		ptLen--
	}

	// give back the frames to the RAM resource manager
	// (so it can go to some other process at some point)
	err := pm.ramRM.Release(frNumbers)
	if err != nil {
		pm.logger.Printf(
			"ERROR releasing frames from process %d: %v",
			p.ProcessNumber, err,
		)
		return err
	}
	pm.logger.Printf(
		"released frames from process %d: %v",
		p.ProcessNumber, frNumbers,
	)

	// remove the corresponing entries from the RAM page table
	// (this is done this way to ensure an entry wasn't missed)
	for _, pn := range pgNumbers {
		delete(p.RAMPageTable, pn)
	}
	pm.logger.Printf(
		"deleted pages from process %d: %v",
		p.ProcessNumber, pgNumbers,
	)

	// make sure we got all the entries!
	// if not, this should panic (becasue it's unexpected)
	if len(p.RAMPageTable) > 0 {
		pm.logger.Panicf(
			"%d page table entries still remain after unloading process %d",
			len(p.RAMPageTable), p.ProcessNumber,
		)
	}

	pm.logger.Printf("process %d has been unloaded", p.ProcessNumber)
	return nil
}

// Reallocate ensures the given process has enough space allocated to it.
// (basically, it handles pages-faults)
func (pm *PageManager) Reallocate(p *process.Process) error {
	err := pm.reallocate(p)
	if err != nil {
		pm.logger.Printf(
			"could not reallocate process %d: %v",
			p.ProcessNumber, err,
		)
		pm.logger.Printf(
			"process %d added to waitlist",
			p.ProcessNumber,
		)
		// the request cannot be granted!
		// add the process to the waitlist
		pm.waitlist = append(pm.waitlist, p)
	} else {
		pm.logger.Printf(
			"process %d reallocated: %v",
			p.ProcessNumber, p.RAMPageTable,
		)
	}
	return err
}

// HandleWaitlist ensures the items in the waitlist are handled eventually.
func (pm *PageManager) HandleWaitlist() {
	pm.logger.Printf("handle waitlist (size: %d)", len(pm.waitlist))

	completed := []int{}
	for i, p := range pm.waitlist {
		err := pm.reallocate(p)
		if err == nil {
			pm.logger.Printf("process %d reallocated", p.ProcessNumber)
			// remove from the waitlist (later)
			completed = append(completed, i)
		} else {
			pm.logger.Printf(
				"ERROR reallocating process %d: %v",
				p.ProcessNumber, err,
			)
		}
	}
	// remove the completed processes from the waitlist
	// this must be done in reverse order to preserve the index integrity
	// (i.e. as you remove values, the indexes shift)
	for i := len(completed) - 1; i >= 0; i-- {
		index := completed[i]
		pm.logger.Printf(
			"process %d removed from waitlist",
			pm.waitlist[index].ProcessNumber,
		)
		// remove the indicated index from the waitlist
		pm.waitlist = append(pm.waitlist[:index], pm.waitlist[index+1:]...)
	}
}

func (pm *PageManager) reallocate(p *process.Process) error {
	pm.logger.Printf("reallocate process %d", p.ProcessNumber)
	state := p.State()

	numFaults := len(state.Faults)
	pm.logger.Printf(
		"claiming %d frames of RAM to reallocate process %d",
		numFaults, p.ProcessNumber,
	)
	frameNums, err := pm.ramRM.Claim(numFaults)
	if err != nil {
		pm.logger.Printf(
			"ERROR claiming frames of RAM to reallocate process %d: %v",
			p.ProcessNumber, err,
		)
		return err
	}
	p.Footprint += numFaults
	pm.logger.Printf(
		"process %d footprint increased by %d to %d",
		p.ProcessNumber, numFaults, p.Footprint,
	)
	i := 0
	for x, v := range state.Faults {
		if !v {
			pm.logger.Panicf(
				"expected true in faults for process %d: %v",
				p.ProcessNumber, state.Faults,
			)
		}
		pn := page.Number(x)
		fn := ivm.FrameNumber(frameNums[i])
		pm.virtualMachine.RAMFrameWrite(fn, ivm.MakeFrame())
		p.RAMPageTable[pn] = fn
		i++
	}
	pm.logger.Printf(
		"process %d RAM page table now %v",
		p.ProcessNumber, p.RAMPageTable,
	)
	state.Faults = ivm.FaultList{}
	p.SetState(state)
	return nil
}

package kernel

import (
  "../vm/ivm"
  "../prog"
  "../config"
  "./page"
  "log"
)

// Kernel houses all the storage and functionality of the OS kernel.
type Kernel struct {
  config config.Config
  virtualMachine ivm.IVM
  // processTable processTableType
  // ramFrameTable frameTableType
  // diskFrameTable frameTableType
}

// MakeKernel makes a kernel with the given virtual machine.
func MakeKernel(virtualMachine ivm.IVM, c config.Config) (Kernel, error) {
  k := Kernel{
    config: c,
    virtualMachine: virtualMachine,
    // processTable: processTableType{},
    // ramFrameTable: frameTableType{},
    // diskFrameTable: frameTableType{},
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

// PageReadRAM reads a page from the RAM at the given page number and page table.
func (k Kernel) PageReadRAM(pageNumber page.Number, pageTable page.Table) page.Page {
	frameNumber := pageTable[pageNumber]
	frame := k.virtualMachine.RAMFrameFetch(frameNumber)
	return page.Page(frame)
}

// PageWriteRAM writes the given page to the RAM at the given page number.
func (k Kernel) PageWriteRAM(page page.Page, pageNumber page.Number, pageTable page.Table) {
	frameNumber := pageTable[pageNumber]
	log.Printf("Writing to Page #%d (Frame #%d) in RAM", pageNumber, frameNumber)
	frame := ivm.Frame(page)
	k.virtualMachine.RAMFrameWrite(frameNumber, frame)
}

// RAMPushPages pushes a given page array into the first available space in RAM.
func (k *Kernel) RAMPushPages(pages []page.Page, pageTable *page.Table) error {
	// TODO: find a good space for the pages to go (and add to the page table)
	// for index, p := range pages {
	// 	pageNumber := page.PageNumber(index)
	// 	// determine a frame number
	// 	// TODO: use a real frame number here (for now, it's just FIFO)
	// 	usedFrameNumbers := k.ramFrameTable.UsedFrameNumbers()
	// 	frameNumber := ivm.FrameNumber(len(usedFrameNumbers))
	// 	if frameNumber >= ivm.RAMNumFrames {
	// 		return nil
	// 	}
	// 	k.ramFrameTable[frameNumber] = pageNumber
	// 	(*pageTable)[pageNumber] = frameNumber
	// 	k.PageWriteRAM(p, pageNumber, *pageTable)
	// }
	return nil
}

// PushProgram pushes a program into the first available space in the VM.
// (this prefers RAM, but falls back to using the disk if necessary; returns error otherwise)
func (k *Kernel) PushProgram(p prog.Program, pageTable *page.Table) error {
	// get the pages for the given program
	// push those pages into the VM and return the result
	// return k.RAMPushPages(PageArrayFromProgram(program), pageTable)
  // TODO: make this work
  return nil
}

// PushOverflowError means there isn't enough storage to hold all provided data.
type PushOverflowError struct{}

func (e PushOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided data."
}

package kernel

import (
	"log"
	"../prog"
	"../vm/ivm"
)

// Page wraps VM Frames to provide a concurrent/continuous view for the process.
type Page ivm.Frame

// PageNumber wraps VM FrameNumbers to provide a location from the page table.
type PageNumber ivm.FrameNumber

// PageTable is a mapping of PageNumbers to VM FrameNumbers
type PageTable map[PageNumber]ivm.FrameNumber

type frameTableType map[ivm.FrameNumber]PageNumber

// UsedFrameNumbers returns the frame numbers already used by the page table.
func (pt PageTable) UsedFrameNumbers() []ivm.FrameNumber {
	frameNumbers := []ivm.FrameNumber{}
	for _, frameNumber := range pt {
		frameNumbers = append(frameNumbers, frameNumber)
	}
	return frameNumbers
}

func (ft frameTableType) UsedFrameNumbers() []ivm.FrameNumber {
	frameNumbers := []ivm.FrameNumber{}
	for frameNumber := range ft {
		frameNumbers = append(frameNumbers, frameNumber)
	}
	return frameNumbers
}

// PageArrayFromFrameArray returns a list of pages for a given list of frames.
func PageArrayFromFrameArray(ary []ivm.Frame) []Page {
	pageArray := make([]Page, len(ary))
	for i, f := range ary {
		pageArray[i] = Page(f)
	}
	return pageArray
}

// PageArrayFromWordArray returns a list of pages for the given list of words.
func PageArrayFromWordArray(ary []ivm.Word) []Page {
	return PageArrayFromFrameArray(ivm.FrameArrayFromWordArray(ary))
}

// PageArrayFromUint32Array returns a list of pages for the given list of uint32's.
func PageArrayFromUint32Array(ary []uint32) []Page {
	return PageArrayFromFrameArray(ivm.FrameArrayFromUint32Array(ary))
}

// PageArrayFromProgram returns a list of pages for the given program.
func PageArrayFromProgram(program prog.Program) []Page {
	// get the words for the given program (uint32 array)
	// organize those words (uint32's) into pages
	return PageArrayFromUint32Array(program.GetWords())
}

// PageReadRAM reads a page from the RAM at the given page number and page table.
func (k Kernel) PageReadRAM(pageNumber PageNumber, pageTable PageTable) Page {
	frameNumber := pageTable[pageNumber]
	frame := k.virtualMachine.RAMFrameFetch(frameNumber)
	return Page(frame)
}

// PageWriteRAM writes the given page to the RAM at the given page number.
func (k Kernel) PageWriteRAM(page Page, pageNumber PageNumber, pageTable PageTable) {
	frameNumber := pageTable[pageNumber]
	log.Printf("Writing to Page #%d (Frame #%d) in RAM", pageNumber, frameNumber)
	frame := ivm.Frame(page)
	k.virtualMachine.RAMFrameWrite(frameNumber, frame)
}

// RAMPushPages pushes a given page array into the first available space in RAM.
func (k *Kernel) RAMPushPages(pages []Page, pageTable *PageTable) error {
	// TODO: find a good space for the pages to go (and add to the page table)
	for index, page := range pages {
		pageNumber := PageNumber(index)
		// determine a frame number
		// TODO: use a real frame number here (for now, it's just FIFO)
		usedFrameNumbers := k.ramFrameTable.UsedFrameNumbers()
		frameNumber := ivm.FrameNumber(len(usedFrameNumbers))
		if frameNumber >= ivm.RAMNumFrames {
			return nil
		}
		k.ramFrameTable[frameNumber] = pageNumber
		(*pageTable)[pageNumber] = frameNumber
		k.PageWriteRAM(page, pageNumber, *pageTable)
	}
	return nil
}

// PushProgram pushes a program into the first available space in the VM.
// (this prefers RAM, but falls back to using the disk if necessary; returns error otherwise)
func (k *Kernel) PushProgram(program prog.Program, pageTable *PageTable) error {
	// get the pages for the given program
	// push those pages into the VM and return the result
	return k.RAMPushPages(PageArrayFromProgram(program), pageTable)
}

// PushOverflowError means there isn't enough storage to hold all provided data.
type PushOverflowError struct{}

func (e PushOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided data."
}

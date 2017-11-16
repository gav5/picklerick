package kernel

import (
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
func PageArrayFromProgram(program prog.Program) ([]Page, error) {
	// get the words for the given program (uint32 array)
	words, err := program.GetWords()
	if err != nil {
		return nil, err
	}
	// organize those words (uint32's) into pages
	return PageArrayFromUint32Array(words), nil
}

// PageReadRAM reads a page from the RAM at the given page number and page table.
func (k Kernel) PageReadRAM(pageNumber PageNumber, pageTable PageTable) Page {
	frameNumber := pageTable[pageNumber]
	frame := k.virtualMachine.RAMFrameFetch(frameNumber)
	return Page(frame)
}

// PageWrite writes the given page to the VM at the given page number.
func (k Kernel) PageWrite(page Page, pageNumber PageNumber, pageTable PageTable) {
	frameNumber := pageTable[pageNumber]
	frame := ivm.Frame(page)
	k.virtualMachine.RAMFrameWrite(frameNumber, frame)
}

// PushPages pushes a given page array into the first available space in the VM.
// (this prefers RAM, but falls back to using disk if necessary; returns error otherwise)
func (k *Kernel) PushPages(pages []Page, pageTable *PageTable) error {
	// TODO: find a good space for the pages to go (and add to the page table)
	for index, page := range pages {
		// TODO: use a real page number here
		k.PageWrite(page, PageNumber(index), *pageTable)
	}
	return nil
}

// PushProgram pushes a program into the first available space in the VM.
// (this prefers RAM, but falls back to using the disk if necessary; returns error otherwise)
func (k *Kernel) PushProgram(program prog.Program, pageTable *PageTable) error {
	// get the pages for the given program
	pages, err := PageArrayFromProgram(program)
	if err != nil {
		return err
	}
	// push those pages into the VM and return the result
	return k.PushPages(pages, pageTable)
}

// PushOverflowError means there isn't enough storage to hold all provided data.
type PushOverflowError struct{}

func (e PushOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided data."
}

package page

import (
	"fmt"

	"../../prog"
	"../../vm/ivm"
)

// Page wraps VM Frames to provide a concurrent/continuous view for the process.
type Page ivm.Frame

// Number wraps VM FrameNumbers to provide a location from the page table.
type Number ivm.FrameNumber

// Table is a mapping of PageNumbers to VM FrameNumbers
type Table map[Number]ivm.FrameNumber

type frameTableType map[ivm.FrameNumber]Number

func (n Number) String() string {
	return fmt.Sprintf("%03X", int(n))
}

// TranslateAddress translates a paged address into a raw frame Address.
func (pt Table) TranslateAddress(addr ivm.Address) ivm.Address {
	pageNumber := Number(addr / ivm.FrameSize)
	frameAddress := addr % ivm.FrameSize
	frameNumber := pt[pageNumber]
	return ivm.Address(frameNumber * ivm.FrameSize) + frameAddress
}

// UsedFrameNumbers returns the frame numbers already used by the page table.
func (pt Table) UsedFrameNumbers() []ivm.FrameNumber {
	frameNumbers := []ivm.FrameNumber{}
	for _, frameNumber := range pt {
		frameNumbers = append(frameNumbers, frameNumber)
	}
	return frameNumbers
}

// FrameNumberForPageNumber returns the frame number for the given page number.
func (pt Table) FrameNumberForPageNumber(pn Number) ivm.FrameNumber {
	return pt[pn]
}

func (ft frameTableType) UsedFrameNumbers() []ivm.FrameNumber {
	frameNumbers := []ivm.FrameNumber{}
	for frameNumber := range ft {
		frameNumbers = append(frameNumbers, frameNumber)
	}
	return frameNumbers
}

// ArrayFromFrameArray returns a list of pages for a given list of frames.
func ArrayFromFrameArray(ary []ivm.Frame) []Page {
	pageArray := make([]Page, len(ary))
	for i, f := range ary {
		pageArray[i] = Page(f)
	}
	return pageArray
}

// ArrayFromWordArray returns a list of pages for the given list of words.
func ArrayFromWordArray(ary []ivm.Word) []Page {
	return ArrayFromFrameArray(ivm.FrameArrayFromWordArray(ary))
}

// ArrayFromUint32Array returns a list of pages for the given list of uint32's.
func ArrayFromUint32Array(ary []uint32) []Page {
	return ArrayFromFrameArray(ivm.FrameArrayFromUint32Array(ary))
}

// ArrayFromProgram returns a list of pages for the given program.
func ArrayFromProgram(program prog.Program) []Page {
	// get the words for the given program (uint32 array)
	// organize those words (uint32's) into pages
	return ArrayFromUint32Array(program.GetWords())
}

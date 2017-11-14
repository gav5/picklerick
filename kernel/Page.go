package kernel

import (
	"../vm/ivm"
)

// Page wraps VM Frames to provide a concurrent/continuous view for the process.
type Page ivm.Frame

// PageNumber wraps VM FrameNumbers to provide a location from the page table.
type PageNumber ivm.FrameNumber

// PageTable is a mapping of PageNumbers to VM FrameNumbers
type PageTable map[PageNumber]ivm.FrameNumber

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

// PageWrite writes the given page to the VM at the given page number.
func PageWrite(vm ivm.IVM, page Page, pageNumber PageNumber) {

}

// PushPages pushes a given page array into the first available space in the VM.
// (this prefers RAM, but falls back to using disk if necessary; returns error otherwise)
func PushPages(vm ivm.IVM, pages []Page) (PageTable, error) {
	// TODO: add to page table
	// TODO: push pages into corresponding frame locations
	// add the given frames to the
	if (int(cur) + len(pages)) < ivm.RAMNumFrames {
		for _, f := range frames {
			vm.RAMFrameWrite(cur, f)
			cur++
		}
	} else if (cur - ivm.RAMNumFrames) < ivm.DiskNumFrames {
		for _, f := range frames {
			vm.DiskFrameWrite(cur-ivm.RAMNumFrames, f)
			cur++
		}
	} else {
		return nil, PushOverflowError{}
	}
	return PageTable{}, nil
}

// PushOverflowError means there isn't enough storage to hold all provided data.
type PushOverflowError struct{}

func (e PushOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided data."
}

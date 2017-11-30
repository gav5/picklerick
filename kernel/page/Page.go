package page

import (
	"fmt"
	"io"
	"log"
	"sort"

	"../../vm/ivm"
)

// Page wraps VM Frames to provide a concurrent/continuous view for the process.
type Page ivm.Frame

// Number wraps VM FrameNumbers to provide a location from the page table.
type Number ivm.FrameNumber

// Table is a mapping of PageNumbers to VM FrameNumbers
type Table map[Number]ivm.FrameNumber

func (n Number) String() string {
	return fmt.Sprintf("%03X", int(n))
}

// TranslateAddress translates a paged address into a raw frame Address.
func (pt Table) TranslateAddress(addr ivm.Address) ivm.Address {
	return ivm.AddressForFramePair(pt.PairForAddress(addr))
}

// PageNumbers returns the page numbers in the page table.
func (pt Table) PageNumbers() []Number {
	outary := make([]Number, len(pt))
	i := 0
	for pn := range pt {
		outary[i] = pn
		i++
	}
	sort.Sort(numbersArray(outary))
	return outary
}

// Pair retuns the paged address pair for the given address
func Pair(addr ivm.Address) (Number, int) {
	pageNumber, index := addr.FramePair()
	return Number(pageNumber), index
}

// PairForAddress returns the frame pair for the given address.
func (pt Table) PairForAddress(addr ivm.Address) (ivm.FrameNumber, int) {
	pageNumber, index := Pair(addr)
	frameNumber := pt.FrameNumberForPageNumber(pageNumber)
	return frameNumber, index
}

// PairForAddressSoft is a little more forgiving than PairForAddress
func (pt Table) PairForAddressSoft(addr ivm.Address) (ivm.FrameNumber, int, bool) {
	pageNumber, index := Pair(addr)
	frameNumber, ok := pt.SoftFrameNumberForPageNumber(pageNumber)
	return frameNumber, index, ok
}

// Pairs returns the page table as parallel arrays.
func (pt Table) Pairs() ([]Number, []ivm.FrameNumber) {
	ptLen := len(pt)
	pgNumbers := make([]Number, ptLen)
	frNumbers := make([]ivm.FrameNumber, ptLen)
	for pn, fn := range pt {
		pgNumbers[ptLen-1] = pn
		frNumbers[ptLen-1] = fn
		ptLen--
	}
	return pgNumbers, frNumbers
}

type numbersArray []Number

func (na numbersArray) Len() int {
	return len(na)
}

func (na numbersArray) Less(i, j int) bool {
	return na[i] < na[j]
}

func (na numbersArray) Swap(i, j int) {
	na[i], na[j] = na[j], na[i]
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
	fn, ok := pt.SoftFrameNumberForPageNumber(pn)
	if !ok {
		log.Panicf(
			"cannot supply a frame number for page %d (page table %v)",
			pn, pt,
		)
	}
	return fn
}

// SoftFrameNumberForPageNumber is a little more forgiving than FrameNumberForPageNumber.
func (pt Table) SoftFrameNumberForPageNumber(pn Number) (ivm.FrameNumber, bool) {
	fn, ok := pt[pn]
	return fn, ok
}

func (pt Table) String() string {
	outval := ""
	pnums := pt.PageNumbers()
	for i, pn := range pnums {
		fn := pt[pn]
		if i > 0 {
			outval += " "
		}
		outval += fmt.Sprintf("%v:%v", pn, fn)
	}
	return outval
}

const fprintColSize = 6

// Fprint prints to the given writer
func (pt Table) Fprint(w io.Writer) error {
	var err error
	pnums := pt.PageNumbers()
	for i, pn := range pnums {
		fn := pt[pn]
		if i%fprintColSize > 0 {
			_, err = fmt.Fprint(w, "  ")
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Fprint(w, "\n")
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(w, "%v:%v", pn, fn)
		if err != nil {
			return err
		}
	}
	return nil
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

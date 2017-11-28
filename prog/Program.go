package prog

import (
	"../vm/ivm"
)

// Program represents the parsed results of a program file
type Program struct {
	Job  Job
	Data Data
}

// GetWords gets the raw binary words for the given program
func (p Program) GetWords() []uint32 {
	outval := make([]uint32, p.binWordSize())

	jobWords := p.Job.getWords()
	dataWords := p.Data.getWords()

	jobBinWordSize := p.Job.binWordSize()
	copy(outval[:jobBinWordSize], jobWords[:])
	copy(outval[jobBinWordSize:], dataWords[:])

	return outval
}

func (p Program) binWordSize() uint8 {
	return p.Job.binWordSize() + p.Data.binWordSize()
}

// NumPages returns the number of pages/frames required to store.
func (p Program) NumPages() uint8 {
	return p.binWordSize() / ivm.FrameSize
}

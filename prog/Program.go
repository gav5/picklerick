package prog

import (
	"encoding/binary"

	"../vm/ivm"
)

var (
	order = binary.BigEndian
)

// Program represents the parsed results of a program file
type Program struct {
	Job  Job
	Data Data
}

// GetWords gets the raw binary words for the given program
func (p Program) GetWords() ([]uint32, error) {
	outval := []uint32{}
	jobWords, jobErr := p.Job.getWords()
	if jobErr != nil {
		return []uint32{}, jobErr
	}
	outval = append(outval, jobWords...)
	dataWords, dataErr := p.Data.getWords()
	if dataErr != nil {
		return []uint32{}, dataErr
	}
	outval = append(outval, dataWords[:]...)
	return outval, nil
}

// SetWords sets the raw binary words for the given program
func (p *Program) SetWords(words []uint32) error {
	var dataWords [45]uint32
	if err := p.Job.setWords(words); err != nil {
		return err
	}
	copy(dataWords[:], words[p.Job.NumberOfWords+1:p.Job.NumberOfWords+46])
	if err := p.Data.setWords(dataWords); err != nil {
		return err
	}
	return nil
}

func (p Program) binWordSize() uint8 {
	return p.Job.binWordSize() + p.Data.binWordSize()
}

// Frames gets the representative frames for the given program.
func (p Program) Frames() ([]ivm.Frame, error) {
	vals, err := p.GetWords()
	if err != nil {
		return nil, err
	}
	return ivm.FrameArrayFromUint32Array(vals), nil
}

// WriteASM writes the assembly instructions to the given file writer
// func (p Program) WriteASM(w io.Writer) error {
// 	return p.Job.WriteASM(w)
// }

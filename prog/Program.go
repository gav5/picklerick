package prog

import (
	"encoding/binary"
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

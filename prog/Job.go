package prog

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"../util"
	"../vm"
)

// Job represents the job section of the program
type Job struct {
	ID             uint8
	NumberOfWords  uint8
	PriorityNumber uint8
	_              uint8
	Instructions   []uint32
}

func (j Job) getWords() ([]uint32, error) {
	outval := make([]uint32, j.binWordSize())
	buf := new(bytes.Buffer)
	header := [4]uint8{
		j.ID,
		j.NumberOfWords,
		j.PriorityNumber,
		0x00,
	}
	if err := binary.Write(buf, order, header); err != nil {
		return []uint32{}, err
	}
	instructions := j.Instructions[:]
	if err := binary.Write(buf, order, instructions); err != nil {
		return []uint32{}, err
	}
	if err := binary.Read(buf, order, &outval); err != nil {
		return []uint32{}, err
	}
	return outval, nil
}

func (j *Job) setWords(val []uint32) error {
	buf := new(bytes.Buffer)
	j.ID = uint8(util.BitExtract32(val[0], 0xff000000))
	j.NumberOfWords = uint8(util.BitExtract32(val[0], 0x00ff0000))
	j.PriorityNumber = uint8(util.BitExtract32(val[0], 0x0000ff00))
	j.Instructions = make([]uint32, j.NumberOfWords)
	if err := binary.Write(buf, order, val[1:j.binWordSize()]); err != nil {
		return err
	}
	if err := binary.Read(buf, order, j.Instructions); err != nil {
		return err
	}
	return nil
}

func (j Job) binWordSize() uint8 {
	return j.NumberOfWords + 1
}

// WriteASM writes the assembly instructions to the given file writer
func (j Job) WriteASM(w io.Writer) error {
	fmt.Fprintf(w, "Job ID: %d\n", j.ID)
	for index, iraw := range j.Instructions {
		instr, err := vm.DecodeInstruction(iraw)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%08X  | %04X |  %s\n", iraw, (index * 4), instr.Assembly())
	}
	return nil
}

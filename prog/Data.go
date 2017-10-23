package prog

import (
	"bytes"
	"encoding/binary"

	"../util"
)

// Data represents the data section of the program
type Data struct {
	InputBufferSize  uint8
	OutputBufferSize uint8
	TempBufferSize   uint8
	_                uint8
	DataBlock        [44]uint32
}

func (d Data) getWords() ([45]uint32, error) {
	outval := [45]uint32{}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, order, d); err != nil {
		return [45]uint32{}, err
	}
	if err := binary.Read(buf, order, &outval); err != nil {
		return [45]uint32{}, err
	}
	return outval, nil
}

func (d *Data) setWords(val [45]uint32) error {
	buf := new(bytes.Buffer)
	d.InputBufferSize = uint8(util.BitExtract32(uint32(val[0]), 0xff000000))
	d.OutputBufferSize = uint8(util.BitExtract32(uint32(val[0]), 0x00ff0000))
	d.TempBufferSize = uint8(util.BitExtract32(uint32(val[0]), 0x0000ff00))
	if err := binary.Write(buf, order, val[1:]); err != nil {
		return err
	}
	if err := binary.Read(buf, order, d.DataBlock[:]); err != nil {
		return err
	}
	return nil
}

func (d Data) binWordSize() uint8 {
	return 45
}

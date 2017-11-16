package prog

// Data represents the data section of the program
type Data struct {
	InputBufferSize  uint8
	OutputBufferSize uint8
	TempBufferSize   uint8
	_                uint8
	DataBlock        [44]uint32
}

func (d Data) getWords() [44]uint32 {
	outval := [44]uint32{}
	copy(outval[:], d.DataBlock[:])
	return outval
}

func (d Data) binWordSize() uint8 {
	return 44
}

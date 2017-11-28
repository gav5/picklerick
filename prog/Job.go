package prog

// Job represents the job section of the program
type Job struct {
	ID             uint8
	NumberOfWords  uint8
	PriorityNumber uint8
	_              uint8
	Instructions   []uint32
}

func (j Job) getWords() []uint32 {
	outval := make([]uint32, j.binWordSize())
	copy(outval[:], j.Instructions[:])
	return outval
}

func (j Job) binWordSize() uint8 {
	return j.NumberOfWords
}

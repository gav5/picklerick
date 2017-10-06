package prog

// Program represents the parsed results of a program file
type Program struct {
	Job  Job
	Data Data
}

// Job represents the job section of the program
type Job struct {
	ID             uint8
	NumberOfWords  uint8
	PriorityNumber uint8
	Instructions   []uint32
}

// Data represents the data section of the program
type Data struct {
	InputBufferSize  uint8
	OutputBufferSize uint8
	TempBufferSize   uint8
	DataBlock        []uint32
}

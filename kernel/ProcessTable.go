package kernel

var procTable map[uint8]Process

func addProcessToProcessTable(process Process) {
	procTable[process.ProcessNumber] = process
}

package kernel

type processTableType map[uint8]Process

func (k *Kernel) addProcessToProcessTable(process Process) {
	k.processTable[process.ProcessNumber] = process
}

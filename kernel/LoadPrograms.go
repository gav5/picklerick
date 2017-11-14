package kernel

import (
	"../prog"
	"../vm/ivm"
)

// LoadPrograms loads the given programs into the virtual machine.
func LoadPrograms(vm ivm.IVM, programs []prog.Program) error {
	cur := ivm.FrameNumber(0)
	var err error
	var words []uint32
	var pageTable PageTable
	for _, p := range programs {
		// add the given pages into the VM (starting with RAM, then disk)
		words, err = p.GetWords()
		if err != nil {
			return err
		}
		pages := PageArrayFromUint32Array(words)
		pageTable, err = PushPages(pages)
		if err != nil {
			return err
		}
		// add the add the program (as a process) to the process table
		procTable[p.Job.ID] = Process{
			CPUID:          0x0,
			ProgramCounter: 0x00,
			CodeSize:       p.Job.NumberOfWords,
			ProcessNumber:  p.Job.ID,
			PageTable:      pageTable,
		}
	}
	return nil
}

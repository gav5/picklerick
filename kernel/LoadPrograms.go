package kernel

import (
	"../prog"
	"../vm/ivm"
)

// LoadPrograms loads the given programs into the virtual machine.
func LoadPrograms(vm ivm.IVM, programs []prog.Program) error {
	for _, p := range programs {
		if err := loadProgram(vm, p); err != nil {
			return err
		}
	}
	return nil
}

func loadProgram(vm ivm.IVM, program prog.Program) error {
	// add the given pages into the VM (starting with RAM, then disk)
	pageTable := PageTable{}
	if err := PushProgram(vm, program, &pageTable); err != nil {
		return err
	}
	// add the add the program (as a process) to the process table
	addProcessToProcessTable(MakeProcess(program, pageTable))
	return nil
}

package kernel

import (
	"../vm/ivm"
)

// ContextSwitch switches context for PCB state to an active state on a CPU core.
func ContextSwitch(vm ivm.IVM, coreNum uint8, process *Process) {
	vm.SetProgramCounter(coreNum, process.ProgramCounter)
	for i, reg := range process.Registers {
		vm.SetRegisterWord(coreNum, ivm.RegisterDesignation(i), reg)
	}
	vm.ResetCore(coreNum)
	// TODO: Move the correct program from disk into RAM
}

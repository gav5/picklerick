package kernel

import (
	"../vm/ivm"
)

// ContextSwitch switches context for PCB state to an active state on a CPU core.
func ContextSwitch(vm ivm.IVM, coreNum uint8, pcb *PCB) {
	vm.SetProgramCounter(coreNum, pcb.ProgramCounter)
	for i, reg := range pcb.Registers {
		vm.SetRegisterWord(coreNum, ivm.RegisterDesignation(i), reg)
	}
	vm.ResetCore(coreNum)
	// TODO: Move the correct program from disk into RAM
}

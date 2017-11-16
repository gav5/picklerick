package kernel

import (
	"../vm/ivm"
)

// ContextSwitch switches context for PCB state to an active state on a CPU core.
func (k *Kernel) ContextSwitch(coreNum uint8, process *Process) {
	k.virtualMachine.SetProgramCounter(coreNum, process.ProgramCounter)
	for i, reg := range process.Registers {
		k.virtualMachine.SetRegisterWord(coreNum, ivm.RegisterDesignation(i), reg)
	}
	k.virtualMachine.ResetCore(coreNum)
	// TODO: Move the correct program from disk into RAM
}

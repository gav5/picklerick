package cpu

import (
	"../proc"
	"../ram"
)

// ContextSwitch switches the context from the frozen PCB state to an active state
func (state *State) ContextSwitch(pcb proc.PCB) {
	copy(state.Registers[:], pcb.Registers[:])
	state.ProgramCounter = pcb.ProgramCounter
	state.Program = pcb.Program
	state.ShouldHalt = false
	ram.ContextSwitch(pcb.Program.Data.DataBlock[:])
}

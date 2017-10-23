package cpu

import (
	"../cpuType"
	"../decoder"
)

// Next runs the currently selected command and moves the program counter to the next one
func (state *State) Next() error {
	// run the currently-selected command
	cmdBinary := state.Program.Job.Instructions[state.ProgramCounter/4]
	instr, err := decoder.InstrFromUint32(cmdBinary)
	if err != nil {
		return err
	}
	base := cpuType.State(*state)
	instr.Exec(&base)
	*state = State(base)
	// move to the next command (if one exists)
	if int(state.ProgramCounter) < len(state.Program.Job.Instructions) {
		// we can still increment the PC counter
		state.ProgramCounter += 4
	} else {
		// time to find another program?
		state.ShouldHalt = true
	}
	return nil
}

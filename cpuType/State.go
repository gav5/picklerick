package cpuType

import (
	"fmt"
	"io"
	"strings"

	"../prog"
	"../ram"
	"../reg"
)

// State represents a given operational state of the CPU
type State struct {
	Registers      reg.List
	ProgramCounter uint32
	// permissions
	// buffers
	// caches
	// activePages

	// NOTE: this is temporary
	// (I had to do this to make this work!)
	Program prog.Program

	// NOTE: should this be here?
	ShouldHalt bool
}

// GetReg fetches the value of the designated register
func (s State) GetReg(d reg.Designation) reg.Storage {
	return s.Registers[d]
}

// SetReg pushes the given value of the designated register
func (s *State) SetReg(d reg.Designation, val reg.Storage) {
	s.Registers[d] = val
}

// GetRegBool fetches the value of the designated register as a boolean
func (s State) GetRegBool(d reg.Designation) bool {
	return s.GetReg(d).GetBool()
}

// SetRegBool pushes the given value of the designated for the given boolean
func (s *State) SetRegBool(d reg.Designation, val bool) {
	s.Registers[d].SetBool(val)
}

// GetAddr fetches the value from the given address
// this also ensures the CPU is currently allowed to access that address
func (s State) GetAddr(a uint32) uint32 {
	// TODO: make sure this is allowed
	// TODO: make sure address is right
	return ram.GetData(a)
}

// SetAddr pushes the value from the given address
// this also ensures the CPU is currently allowed to access that address
func (s State) SetAddr(a uint32, d uint32) {
	// TODO: make sure this is allowed
	// TODO: make sure address is right
	ram.SetData(a, d)
}

// GetPC gets the program counter value from the CPU state
func (s State) GetPC() uint32 {
	return s.ProgramCounter
}

// SetPC sets the program counter value for the CPU state
func (s *State) SetPC(val uint32) {
	s.ProgramCounter = val
}

// Halt halts the currently-running program
// (the system will either go to the next program or just exit)
func (s *State) Halt() {
	s.ShouldHalt = true
}

// Write sends the state to the given writer
func (s State) Write(w io.Writer) error {
	bars := strings.Repeat("=", 13)
	fmt.Fprintln(w, bars)
	if _, err := fmt.Fprintf(w, "pc:  %v\n", s.ProgramCounter); err != nil {
		return err
	}
	fmt.Fprintln(w, bars)
	for i, reg := range s.Registers {
		if _, err := fmt.Fprintf(w, "r%02d: %v\n", i, reg); err != nil {
			return err
		}
	}
	return nil
}

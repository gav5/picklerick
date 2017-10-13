package cpu

import (
	"fmt"
	"io"
	"strings"

	"../ram"
	"../reg"
)

// State represents a given operational state of the CPU
type State struct {
	registers      reg.List
	programCounter reg.Storage
	// permissions
	// buffers
	// caches
	// activePages
}

// TODO: describe context-switching (should take a PCB and apply it)

// GetReg fetches the value of the designated register
func (s State) GetReg(d reg.Designation) reg.Storage {
	return s.registers[d]
}

// SetReg pushes the given value of the designated register
func (s *State) SetReg(d reg.Designation, val reg.Storage) {
	s.registers[d] = val
}

// GetAddr fetches the value from the given address
// this also ensures the CPU is currently allowed to access that address
func (s State) GetAddr(a ram.Address) ram.Word {
	// TODO: make sure this is allowed
	// TODO: make sure address is right
	return ram.GetData(a)
}

// SetAddr pushes the value from the given address
// this also ensures the CPU is currently allowed to access that address
func (s State) SetAddr(a ram.Address, d ram.Word) {
	// TODO: make sure this is allowed
	// TODO: make sure address is right
	ram.SetData(a, d)
}

// Write sends the state to the given writer
func (s State) Write(w io.Writer) error {
	bars := strings.Repeat("=", 13)
	fmt.Fprintln(w, bars)
	if _, err := fmt.Fprintf(w, "pc:  %v\n", s.programCounter); err != nil {
		return err
	}
	fmt.Fprintln(w, bars)
	for i, reg := range s.registers {
		if _, err := fmt.Fprintf(w, "r%02d: %v\n", i, reg); err != nil {
			return err
		}
	}
	return nil
}

package ivm

// State represents the operational state of a CPU Core.
type State struct {
  ProgramCounter Address
  Registers [NumCoreRegisters]Word
  Halt bool
  Caches []Frame
  Faults []FrameNumber
  Error error
}

// MakeState builds a new blank state.
func MakeState() State {
  s := State{
    ProgramCounter: 0x00000000,
    Registers: [NumCoreRegisters]Word{},
    Halt: false,
    Caches: []Frame{},
    Faults: []FrameNumber{},
    Error: nil,
  }
  // zero out each register
  for i := range s.Registers {
    s.Registers[i] = Word(0x00000000)
  }
  return s
}

// Next builds an initial next state from an existing state.
func (s State) Next() State {
  next := State{
    ProgramCounter: s.ProgramCounter + 4,
    Registers: [NumCoreRegisters]Word{},
    Halt: false,
    Caches: []Frame{},
    Faults: []FrameNumber{},
    Error: nil,
  }
  copy(next.Registers[:], s.Registers[:])
  copy(next.Caches[:][:], s.Caches[:][:])
  copy(next.Faults[:], s.Faults[:])
  return next
}

// RegisterWord returns the given register word value.
func (s State) RegisterWord(regNum RegisterDesignation) Word {
	return s.Registers[regNum]
}

// SetRegisterWord sets the given register to the given word value.
func (s *State) SetRegisterWord(regNum RegisterDesignation, val Word) {
	s.Registers[regNum] = val
}

// RegisterUint32 returns the given register uint32 value.
func (s State) RegisterUint32(regNum RegisterDesignation) uint32 {
	return s.RegisterWord(regNum).Uint32()
}

// SetRegisterUint32 sets the given register to the given uint32 value.
func (s *State) SetRegisterUint32(regNum RegisterDesignation, val uint32) {
	s.SetRegisterWord(regNum, WordFromUint32(val))
}

// RegisterInt32 returns the given register int32 value.
func (s State) RegisterInt32(regNum RegisterDesignation) int32 {
	return s.RegisterWord(regNum).Int32()
}

// SetRegisterInt32 sets the given register to the given int32 value.
func (s *State) SetRegisterInt32(regNum RegisterDesignation, val int32) {
	s.SetRegisterWord(regNum, WordFromInt32(val))
}

// RegisterBool returns the given register bool value.
func (s State) RegisterBool(regNum RegisterDesignation) bool {
	return s.RegisterWord(regNum).Bool()
}

// SetRegisterBool sets the given register to the given bool value.
func (s *State) SetRegisterBool(regNum RegisterDesignation, val bool) {
	s.SetRegisterWord(regNum, WordFromBool(val))
}

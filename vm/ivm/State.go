package ivm

// State represents the operational state of a CPU Core.
type State struct {
  ProgramCounter Address
  Registers [NumCoreRegisters]Word
  Halt bool
  Caches FrameCache
  Faults FaultList
  Error error
  isSleep bool
}

// MakeState builds a new blank state.
func MakeState() State {
  s := State{
    ProgramCounter: 0x00000000,
    Registers: [NumCoreRegisters]Word{},
    Halt: false,
    Caches: FrameCache{},
    Faults: map[FrameNumber]bool{},
    Error: nil,
    isSleep: false,
  }
  // zero out each register
  for i := range s.Registers {
    s.Registers[i] = Word(0x00000000)
  }
  return s
}

// Sleep makes a state for the sleep process.
func Sleep() State {
  return State{
    ProgramCounter: 0x00000000,
    Registers: [NumCoreRegisters]Word{},
    Halt: false,
    Caches: FrameCache{
      0x00: Frame{0x13000000},
    },
    Faults: map[FrameNumber]bool{},
    Error: nil,
    isSleep: true,
  }
}

// Next builds an initial next state from an existing state.
func (s State) Next() State {
  if s.isSleep {
    // sleep states produce the same value
    // (because duh!)
    return s
  }
  next := State{
    ProgramCounter: s.ProgramCounter + 4,
    Registers: [NumCoreRegisters]Word{},
    Halt: false,
    Caches: s.Caches.Copy(),
    Faults: s.Faults.Copy(),
    Error: nil,
  }
  copy(next.Registers[:], s.Registers[:])
  return next
}

// Apply considers a new state that should replace the existing one.
func (s *State) Apply(n State) State {
  s.ProgramCounter = n.ProgramCounter
  copy(s.Registers[:], n.Registers[:])
  s.Halt = false
  s.Caches = n.Caches.Copy()
  s.Faults = n.Faults.Copy()
  s.Error = nil
  return *s
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

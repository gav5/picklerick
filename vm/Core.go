package vm

import (
	"./ivm"
)

// Core describes a CPU Core in the virtual machine.
type Core struct {
	PC         ivm.Address
	Registers  [ivm.NumCoreRegisters]ivm.Word
	ShouldHalt bool
}

// ProgramCounter returns the value of the program counter (PC).
func (c Core) ProgramCounter() ivm.Address {
	return c.PC
}

// SetProgramCounter sets the value of the program counter to the given value.
func (c *Core) SetProgramCounter(val ivm.Address) {
	c.PC = val
}

// Halt halts current program execution.
func (c *Core) Halt() {
	c.ShouldHalt = true
}

// Reset resets the state to start over (undoing a halt)
func (c *Core) Reset() {
	c.ShouldHalt = false
}

// RegisterWord returns the given register word value.
func (c Core) RegisterWord(regNum ivm.RegisterDesignation) ivm.Word {
	return c.Registers[regNum]
}

// SetRegisterWord sets the given register to the given word value.
func (c *Core) SetRegisterWord(regNum ivm.RegisterDesignation, val ivm.Word) {
	c.Registers[regNum] = val
}

// RegisterUint32 returns the given register uint32 value.
func (c Core) RegisterUint32(regNum ivm.RegisterDesignation) uint32 {
	return c.RegisterWord(regNum).Uint32()
}

// SetRegisterUint32 sets the given register to the given uint32 value.
func (c *Core) SetRegisterUint32(regNum ivm.RegisterDesignation, val uint32) {
	c.SetRegisterWord(regNum, ivm.WordFromUint32(val))
}

// RegisterInt32 returns the given register int32 value.
func (c Core) RegisterInt32(regNum ivm.RegisterDesignation) int32 {
	return c.RegisterWord(regNum).Int32()
}

// SetRegisterInt32 sets the given register to the given int32 value.
func (c *Core) SetRegisterInt32(regNum ivm.RegisterDesignation, val int32) {
	c.SetRegisterWord(regNum, ivm.WordFromInt32(val))
}

// RegisterBool returns the given register bool value.
func (c Core) RegisterBool(regNum ivm.RegisterDesignation) bool {
	return c.RegisterWord(regNum).Bool()
}

// SetRegisterBool sets the given register to the given bool value.
func (c *Core) SetRegisterBool(regNum ivm.RegisterDesignation, val bool) {
	c.SetRegisterWord(regNum, ivm.WordFromBool(val))
}

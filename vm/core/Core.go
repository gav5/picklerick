package core

import (
	"fmt"
	"log"
	"../ivm"
	"../decoder"
	"../../kernel/process"
)

// Core describes a CPU Core in the virtual machine.
type Core struct {
	CoreNum    uint8
	PC         ivm.Address
	Registers  [ivm.NumCoreRegisters]ivm.Word
	ShouldHalt bool
	currentContext *Context
}

// New makes a new core.
func New(coreNum uint8) *Core {
	core := &Core{
		CoreNum:    coreNum,
		PC:         0x00000000,
		Registers:  [ivm.NumCoreRegisters]ivm.Word{},
		ShouldHalt: false,
	}
	return core
}

// Call runs the instruction at PC, increments PC (unless manually set).
func (c *Core) Call() Signal {
	callsign := fmt.Sprintf("[CORE%d:%v]", c.CoreNum, c.PC)
	log.Printf("%s Begin execution\n", callsign)
	if c.currentContext == nil {
		log.Printf("%s ERR NO CONTEXT\n", callsign)
		return Signal{CoreNum: c.CoreNum, Error: NoContextError{}, Halted: false}
	}
	instructionRAW := c.currentContext.VM.RAMAddressFetchUint32(c.PC)
	log.Printf("%s InstructionRAW: %08X\n", callsign, instructionRAW)
	instruction, err := decoder.DecodeInstruction(instructionRAW)
	if err != nil {
		log.Printf("%s INSTR DECODE ERR: %v\n", callsign, err)
		return Signal{CoreNum: c.CoreNum, Error: err, Halted: false}
	}
	log.Printf("%s Decoded to: %s\n", callsign, instruction.Assembly())
	ip := c.currentContext.VM.InstructionProxy(c)
	log.Printf("%s Executing instruction...\n", callsign)
	instruction.Execute(ip)
	log.Printf("%s Instruction executed!\n", callsign)
	if c.ShouldHalt {
		log.Printf("%s HALTED!\n", callsign)
		return Signal{CoreNum: c.CoreNum, Error: nil, Halted: true}
	}
	defer func() {
		c.PC += 4
	}()
	return Signal{CoreNum: c.CoreNum, Error: nil, Halted: false}
}

// Apply a context to the given CPU Core.
func (c *Core) Apply(context *Context) {
	c.ShouldHalt = false
	c.PC = context.NextProcess.ProgramCounter
	context.NextProcess.CPUID = c.CoreNum
	copy(c.Registers[:], context.NextProcess.Registers[:])
	c.currentContext = context
}

// Save applies the CPU Core's current state to the process.
func (c *Core) Save() {
	c.currentContext.NextProcess.ProgramCounter = c.PC
	copy(c.currentContext.NextProcess.Registers[:], c.Registers[:])
	if c.ShouldHalt {
		c.currentContext.NextProcess.Status = process.Done
	}
}

// Run runs the core.
// func (c *Core) Run(vm ivm.IVM) *Endpoint {
// 	go func() {
// 		c.PC = 0x00000000
// 		for {
// 			// wait until the VM says to go
// 			tickSignal := <- c.tickChannel
//
// 			// stop if the VM says to do so
// 			// (otherwise, keep going)
// 			if tickSignal.ShouldStop {
// 				close(c.tockChannel)
// 				return
// 			}
//
// 			instructionRAW := vm.RAMAddressFetchUint32(c.PC)
// 			instruction, err := decoder.DecodeInstruction(instructionRAW)
// 			if err != nil {
// 				log.Printf("[%d] ERR: %v\n", c.CoreNum, err)
// 				c.tockChannel <- Tock{Error: err, Halted: false}
// 				break
// 			}
// 			ip := vm.InstructionProxy(c)
// 			instruction.Execute(ip)
// 			log.Printf("[%d] %s\n", c.CoreNum, instruction.Assembly())
// 			if c.ShouldHalt {
// 				log.Printf("[%d] HALTING...\n", c.CoreNum)
// 				c.tockChannel <- Tock{Error: nil, Halted: true}
// 				break
// 			} else {
// 				c.tockChannel <- Tock{Error: nil, Halted: false}
// 			}
// 			c.PC += 4
// 		}
// 	}()
// 	return &Endpoint{
// 		TickSender: (chan<- Tick)(c.tickChannel),
// 		TockReceiver: (<-chan Tock)(c.tockChannel),
// 	}
// }

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

package core

import (
	"fmt"
	"log"

	"../../kernel/process"
	"../../util/logger"
	"../decoder"
	"../ivm"
)

// Core describes a CPU Core in the virtual machine.
type Core struct {
	CoreNum   uint8
	Process   process.Process
	Next      ivm.State
	logger    *log.Logger
	snapshots []Snapshot
}

// Make builds a new core.
func Make(coreNum uint8) Core {
	return Core{
		CoreNum:   coreNum,
		Process:   process.Sleep(), // <- obviously this is a bad assumption
		Next:      ivm.MakeState(),
		logger:    logger.New(fmt.Sprintf("core%d", coreNum)),
		snapshots: []Snapshot{},
	}
}

// MakeArray builds an array of cores.
// (the size is determined by the IVM number of cores)
func MakeArray() [ivm.NumCores]Core {
	cores := [ivm.NumCores]Core{}
	for i := range cores {
		cores[i] = Make(uint8(i + 1))
	}
	return cores
}

// Mock builds a core for testing.
func Mock(sampleProcess process.Process) Core {
	return Core{
		CoreNum:   0,
		Process:   sampleProcess,
		Next:      sampleProcess.State().Next(),
		logger:    logger.Dummy(),
		snapshots: []Snapshot{},
	}
}

// Call runs the instruction at PC, increments PC (unless manually set).
func (c *Core) Call() {
	c.logger.SetPrefix(c.logPrefix())

	// get the current instruction
	instruction, err := c.currentInstruction()
	if err != nil {
		c.logger.Printf("INSTR DECODE ERR: %v", err)
		c.Next.Error = err
		return
	}

	// execute the current instruction
	// (passing in the next state of the system)
	// (this will be handled by the virtual machine later)
	c.logger.Printf("Executing %v", instruction.Assembly())
	ip := ivm.MakeInstructionProxy(&c.Next)
	instruction.Execute(ip)
}

func (c Core) currentInstruction() (ivm.Instruction, error) {
	s := c.Process.State()
	pc := s.ProgramCounter
	raw := s.Caches.AddressFetchWord(pc).Uint32()
	instr, err := decoder.DecodeInstruction(raw)
	if err != nil {
		return nil, err
	}
	return instr, nil
}

func (c Core) logPrefix() string {
	if c.Process.IsSleep() {
		return fmt.Sprintf(
			"core%d:00/zzzz | ",
			c.CoreNum,
		)
	}
	return fmt.Sprintf(
		"core%d:%02d/%04x | ",
		c.CoreNum, c.Process.ProcessNumber,
		uint(c.Process.State().ProgramCounter),
	)
}

package core

import (
	"fmt"
	"log"
	"../ivm"
	"../decoder"
	"../../util/logger"
	"../../kernel/process"
)

// Core describes a CPU Core in the virtual machine.
type Core struct {
	CoreNum	uint8
	Process process.Process
	Next ivm.State
	logger *log.Logger
}

// Make builds a new core.
func Make(coreNum uint8) Core {
	return Core{
		CoreNum: coreNum,
		Process: process.Sleep(), // <- obviously this is a bad assumption
		Next: ivm.MakeState(),
		logger: logger.New(fmt.Sprintf("core%d", coreNum)),

	}
}

// MakeArray builds an array of cores.
// (the size is determined by the IVM number of cores)
func MakeArray() [ivm.NumCores]Core {
	cores := [ivm.NumCores]Core{}
	for i := range cores {
		cores[i] = Make(uint8(i+1))
	}
	return cores
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
	pc := c.Process.State.ProgramCounter
	raw := c.Process.State.Caches.AddressFetchWord(pc).Uint32()
	instr, err := decoder.DecodeInstruction(raw)
	if err != nil {
		return nil, err
	}
	return instr, nil
}

// Apply a process to the given CPU Core.
func (c *Core) Apply(p *process.Process) {
	if p == nil {
		c.logger.Printf("NO JOB")

		c.Process = process.Sleep()
		// c.ShouldHalt = true
		return
	}
	// if c.CurrentProcess != p {
	// 	c.logger.Printf(
	// 		"[CPU%d] Job #%d\n",
	// 		c.CoreNum, p.ProcessNumber,
	// 	)
	// 	c.ShouldHalt = false
	// 	c.PC = p.ProgramCounter
	// 	(*p).CPUID = c.CoreNum
	// 	copy(c.Registers[:], p.Registers[:])
	// 	c.CurrentProcess = p
	// }
}

// Save applies the CPU Core's current state to the process.
func (c *Core) Save() {
	// if c.CurrentProcess == nil {
	// 	return
	// }
	// c.CurrentProcess.ProgramCounter = c.PC
	// copy(c.CurrentProcess.Registers[:], c.Registers[:])
	// if c.ShouldHalt {
	// 	c.CurrentProcess.Status = process.Done
	// }
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
		uint(c.Process.State.ProgramCounter),
	)
}

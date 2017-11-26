package core

import (
	"fmt"
	"log"
	"../decoder"
	"../../kernel/process"
	"../ivm"
	// "../../kernel/page"
)

// Core describes a CPU Core in the virtual machine.
type Core struct {
	CoreNum	uint8
	Process *process.Process
	Next ivm.State
}

// New makes a new core.
func New(coreNum uint8) *Core {
	core := &Core{
		CoreNum: coreNum,
		Next: ivm.MakeState(),
	}
	return core
}

// Call runs the instruction at PC, increments PC (unless manually set).
func (c *Core) Call() {
	callsign := fmt.Sprintf(
		"[CORE%d:%04x]",
		c.CoreNum, uint(c.Process.State.ProgramCounter),
	)
	log.Printf("%s Begin execution\n", callsign)

	if c.Process == nil {
		log.Printf("%s NO PROCESS\n", callsign)
		// probably just nearing the end, so do nothing!
		return
	}

	// get the current instruction
	instruction, err := c.currentInstruction()
	if err != nil {
		log.Printf("%s INSTR DECODE ERR: %v\n", callsign, err)
		c.Next.Error = err
		return
	}

	// execute the current instruction
	// (passing in the next state of the system)
	// (this will be handled by the virtual machine later)
	ip := ivm.MakeInstructionProxy(&c.Next)
	instruction.Execute(ip)
}

func (c Core) currentInstruction() (ivm.Instruction, error) {
	instructions := c.Process.State.Caches[:c.Process.Program.InputBufferSize]
	frameNum := int(c.Process.State.ProgramCounter/ivm.FrameSize)
	frameIndex := int(c.Process.State.ProgramCounter%ivm.FrameSize)
	raw := uint32(instructions[frameNum][frameIndex])
	instr, err := decoder.DecodeInstruction(raw)
	if err != nil {
		return nil, err
	}
	return instr, nil
}

// Apply a process to the given CPU Core.
func (c *Core) Apply(p *process.Process) {
	if p == nil {
		log.Printf("[CPU%d] NO JOB\n", c.CoreNum)

		c.Process = nil
		// c.ShouldHalt = true
		return
	}
	// if c.CurrentProcess != p {
	// 	log.Printf(
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

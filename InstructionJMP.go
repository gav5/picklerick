package main

import "fmt"

// InstructionJMP jumps to a specified location
type InstructionJMP struct {
	args InstructionArgsJump
}

func (iJMP InstructionJMP) exec(pcb PCB) PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (iJMP InstructionJMP) ASM() string {
	return fmt.Sprintf("JMP %s", iJMP.args.ASM())
}

func makeInstructionJMP(args InstructionArgs) InstructionBase {
	return InstructionJMP{args: args.jumpFormat()}
}

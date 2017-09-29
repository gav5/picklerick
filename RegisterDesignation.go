package main

import "fmt"

// RegisterDesignation references a register in the CPU
type RegisterDesignation uint8

// ASM returns the representation in assembly language
func (rd RegisterDesignation) ASM() string {
	return fmt.Sprintf("R%d", uint8(rd))
}

const (
	// R0 register
	R0 RegisterDesignation = 0
	// R1 register
	R1 RegisterDesignation = 1
	// R2 register
	R2 RegisterDesignation = 2
	// R3 register
	R3 RegisterDesignation = 3
	// R4 register
	R4 RegisterDesignation = 4
	// R5 register
	R5 RegisterDesignation = 5
	// R6 register
	R6 RegisterDesignation = 6
	// R7 register
	R7 RegisterDesignation = 7
	// R8 register
	R8 RegisterDesignation = 8
	// R9 register
	R9 RegisterDesignation = 9
	// R10 register
	R10 RegisterDesignation = 10
	// R11 register
	R11 RegisterDesignation = 11
	// R12 register
	R12 RegisterDesignation = 12
	// R13 register
	R13 RegisterDesignation = 13
	// R14 register
	R14 RegisterDesignation = 14
	// R15 register
	R15 RegisterDesignation = 15
)

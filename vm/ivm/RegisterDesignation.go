package ivm

import "fmt"

// RegisterDesignation references a register in the CPU
type RegisterDesignation uint8

func (rd RegisterDesignation) String() string {
	return rd.ASM()
}

// ASM returns the representation in assembly language
func (rd RegisterDesignation) ASM() string {
	return fmt.Sprintf("R%d", uint8(rd))
}

const (
	// R0 is register 0
	R0 RegisterDesignation = iota
	// R1 is register 1
	R1
	// R2 is register 2
	R2
	// R3 is register 3
	R3
	// R4 is register 4
	R4
	// R5 is register 5
	R5
	// R6 is register 6
	R6
	// R7 is register 7
	R7
	// R8 is register 8
	R8
	// R9 is register 9
	R9
	// R10 is register 10
	R10
	// R11 is register 11
	R11
	// R12 is register 12
	R12
	// R13 is register 13
	R13
	// R14 is register 14
	R14
	// R15 is register 15
	R15
)

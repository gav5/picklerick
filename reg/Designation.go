package reg

import "fmt"

// Designation references a register in the CPU
type Designation uint8

// ASM returns the representation in assembly language
func (rd Designation) ASM() string {
	return fmt.Sprintf("R%d", uint8(rd))
}

const (
	// R0 register
	R0 Designation = 0
	// R1 register
	R1 Designation = 1
	// R2 register
	R2 Designation = 2
	// R3 register
	R3 Designation = 3
	// R4 register
	R4 Designation = 4
	// R5 register
	R5 Designation = 5
	// R6 register
	R6 Designation = 6
	// R7 register
	R7 Designation = 7
	// R8 register
	R8 Designation = 8
	// R9 register
	R9 Designation = 9
	// R10 register
	R10 Designation = 10
	// R11 register
	R11 Designation = 11
	// R12 register
	R12 Designation = 12
	// R13 register
	R13 Designation = 13
	// R14 register
	R14 Designation = 14
	// R15 register
	R15 Designation = 15
)

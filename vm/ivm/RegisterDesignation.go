package ivm

import "fmt"

// RegisterDesignation references a register in the CPU
type RegisterDesignation uint8

// ASM returns the representation in assembly language
func (rd RegisterDesignation) ASM() string {
	return fmt.Sprintf("R%d", uint8(rd))
}

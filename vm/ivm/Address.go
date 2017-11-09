package ivm

import (
	"fmt"
)

// Address encapsulates an address used to fetch from the bus
type Address uint32

// Hex returns the hexidecimal (base-16) string representation
func (addr Address) Hex() string {
	return fmt.Sprintf("0x%08X", uint32(addr))
}

// Dec returns the decimal (base-10) string representation
func (addr Address) Dec() string {
	return fmt.Sprintf("%d", uint32(addr))
}

func (addr Address) String() string {
	return fmt.Sprintf("%08x", uint32(addr))
}

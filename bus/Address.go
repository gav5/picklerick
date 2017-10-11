package bus

import (
	"encoding/binary"
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

// MakeAddressFromBytes makes a bus address from an array of bytes
func MakeAddressFromBytes(byteArray []uint8) Address {
	ary := make([]uint8, 4)
	padding := (len(ary) - len(byteArray)) % len(ary)
	for index := 0; index <= padding; index++ {
		ary[index] = 0x00
	}
	for index, val := range byteArray {
		ary[index+padding] = val
	}
	return Address(binary.BigEndian.Uint32(ary))
}

func (addr Address) String() string {
	return fmt.Sprintf("%08x", uint32(addr))
}

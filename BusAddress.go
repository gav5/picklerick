package main

import (
	"encoding/binary"
	"fmt"
)

// BusAddress encapsulates an address used to fetch from the bus
type BusAddress uint32

// Hex returns the hexidecimal (base-16) string representation
func (addr BusAddress) Hex() string {
	return fmt.Sprintf("%#08x", uint32(addr))
}

// Dec returns the decimal (base-10) string representation
func (addr BusAddress) Dec() string {
	return fmt.Sprintf("%d", uint32(addr))
}

func makeBusAddressFromBytes(byteArray []uint8) BusAddress {
	ary := make([]uint8, 4)
	padding := (len(ary) - len(byteArray)) % len(ary)
	for index := 0; index <= padding; index++ {
		ary[index] = 0x00
	}
	for index, val := range byteArray {
		ary[index+padding] = val
	}
	return BusAddress(binary.BigEndian.Uint32(ary))
}

func (addr BusAddress) String() string {
	return fmt.Sprintf("%08x", uint32(addr))
}

package ivm

import (
	"fmt"
)

// Address encapsulates an address used to fetch from the bus
type Address uint32

// Hex returns the hexidecimal (base-16) string representation
func (addr Address) Hex() string {
	return fmt.Sprintf("0x%03X", uint32(addr))
}

// Dec returns the decimal (base-10) string representation
func (addr Address) Dec() string {
	return fmt.Sprintf("%d", uint32(addr))
}

func (addr Address) String() string {
	return fmt.Sprintf("%03x", uint32(addr))
}

// FramePair returns the frame number and index in the frame.
func (addr Address) FramePair() (FrameNumber, int) {
	index := addr.Index()
	return FrameNumber(index / FrameSize), int(index % FrameSize)
}

// AddressForIndex returns an address for an incrementing array index.
func AddressForIndex(index int) Address {
	return Address(index * 4)
}

// Index returns the array index for the given address.
func (addr Address) Index() int {
	return int(addr / 4)
}

// AddressForFramePair returns the address for the given frame pair.
func AddressForFramePair(fn FrameNumber, index int) Address {
	return AddressForIndex(int(fn*FrameSize) + index)
}

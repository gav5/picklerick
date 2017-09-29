package main

import "math/bits"

// BitExtract8 extracts masked bits from an unsigned 8-bit value
func BitExtract8(i uint8, m uint8) uint8 {
	shamt := uint(bits.TrailingZeros8(m))
	return (i & m) >> shamt
}

// BitExtract16 extracts masked bits from an unsigned 16-bit value
func BitExtract16(i uint16, m uint16) uint16 {
	shamt := uint(bits.TrailingZeros16(m))
	return (i & m) >> shamt
}

// BitExtract32 extracts masked bits from an unsigned 32-bit value
func BitExtract32(i uint32, m uint32) uint32 {
	shamt := uint(bits.TrailingZeros32(m))
	return (i & m) >> shamt
}

// BitExtract64 extracts masked bits from an unsigned 64-bit value
func BitExtract64(i uint64, m uint64) uint64 {
	shamt := uint(bits.TrailingZeros64(m))
	return (i & m) >> shamt
}

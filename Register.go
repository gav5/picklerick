package main

import "fmt"

// Register - represents a single register on the CPU
type Register uint32

func (r Register) String() string {
	return fmt.Sprintf("%08X", uint32(r))
}

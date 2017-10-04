package reg

import "fmt"

// Storage represents the storage of a single 32-bit register on the CPU
type Storage uint32

func (s Storage) String() string {
	return fmt.Sprintf("%08X", uint32(s))
}

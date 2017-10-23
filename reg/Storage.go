package reg

import "fmt"

// Storage represents the storage of a single 32-bit register on the CPU
type Storage uint32

func (s Storage) String() string {
	return fmt.Sprintf("%08X", uint32(s))
}

// GetBool gives you the boolean value of the register storage
func (s Storage) GetBool() bool {
	return (s > 0)
}

// SetBool sets the boolean value of the register storage
func (s *Storage) SetBool(val bool) {
	if val {
		*s = 1
	} else {
		*s = 0
	}
}

package ivm

import "fmt"

// Word represents a word stored on the virtual machine.
// This can be used on a register, in RAM, or on disk.
type Word uint32

// Uint32 returns the uint32 value of the word.
func (w Word) Uint32() uint32 {
	return uint32(w)
}

// Int32 returns the int32 value of the word.
func (w Word) Int32() int32 {
	return int32(w)
}

// Bool returns the boolean value of the word.
func (w Word) Bool() bool {
	return w > 0
}

// WordFromUint32 returns a word for the given uint32 value.
func WordFromUint32(val uint32) Word {
	return Word(val)
}

// WordFromInt32 returns a word for the given int32 value.
func WordFromInt32(val int32) Word {
	return Word(val)
}

// WordFromBool returns a word for the given boolean value.
func WordFromBool(val bool) Word {
	if val {
		return Word(0x1)
	}
	return Word(0x0)
}

func (w Word) String() string {
	return fmt.Sprintf("0x%08X", uint32(w))
}

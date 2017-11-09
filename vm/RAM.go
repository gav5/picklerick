package vm

import (
	"fmt"
	"io"
	"os"
	"strings"

	"./ivm"
)

// RAMNumWords is the number of words in the virtual machine RAM.
const RAMNumWords = 1024

// FrameSize is the size of a frame in RAM.
const FrameSize = 16

// RAM describes the virtual machine's RAM module.
type RAM [RAMNumWords]ivm.Word

// AddressFetchWord returns the word value at the given address.
func (r RAM) AddressFetchWord(addr ivm.Address) ivm.Word {
	return r[addr]
}

// AddressWriteWord writes the given word value to the given address.
func (r *RAM) AddressWriteWord(addr ivm.Address, val ivm.Word) {
	r[addr] = val
}

// AddressFetchUint32 returns the uint32 value at the given address.
func (r RAM) AddressFetchUint32(addr ivm.Address) uint32 {
	return r.AddressFetchWord(addr).Uint32()
}

// AddressWriteUint32 writes the given uint32 value to the given address.
func (r *RAM) AddressWriteUint32(addr ivm.Address, val uint32) {
	r.AddressWriteWord(addr, ivm.WordFromUint32(val))
}

// AddressFetchInt32 returns the int32 value at the given address.
func (r RAM) AddressFetchInt32(addr ivm.Address) int32 {
	return r.AddressFetchWord(addr).Int32()
}

// AddressWriteInt32 writes the given int32 value to the given address.
func (r *RAM) AddressWriteInt32(addr ivm.Address, val int32) {
	r.AddressWriteWord(addr, ivm.WordFromInt32(val))
}

// AddressFetchBool returns the bool value at the given address.
func (r RAM) AddressFetchBool(addr ivm.Address) bool {
	return r.AddressFetchWord(addr).Bool()
}

// AddressWriteBool writes the given bool value to the given address.
func (r *RAM) AddressWriteBool(addr ivm.Address, val bool) {
	r.AddressWriteWord(addr, ivm.WordFromBool(val))
}

// Print prints the contents of RAM to Stdout
func (r RAM) Print() error {
	return r.Fprint(os.Stdout)
}

// Fprint prints the contents of RAM to the given writer
func (r RAM) Fprint(w io.Writer) error {
	const numcolumns = 8
	border := strings.Repeat("-", 11*numcolumns+3)
	var err error
	fmt.Fprint(w, "\n")
	for index, val := range r {
		if (index % FrameSize) == 0 {
			if index > 0 {
				if _, err = fmt.Fprintln(w, border); err != nil {
					return err
				}
			}
			if _, err = fmt.Fprintf(w, "Frame %02X\n", index/FrameSize); err != nil {
				return err
			}
		}
		switch index % numcolumns {
		case 0:
			if _, err = fmt.Fprintf(w, "%04X: %08X", index, val); err != nil {
				return err
			}
		case (numcolumns - 1):
			if _, err = fmt.Fprintf(w, " | %08X\n", val); err != nil {
				return err
			}
		default:
			if _, err = fmt.Fprintf(w, " | %08X", val); err != nil {
				return err
			}
		}
	}
	return nil
}

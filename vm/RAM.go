package vm

import (
	"fmt"
	"io"
	"os"
	// "strings"

	"./ivm"
)

// RAM describes the virtual machine's RAM module.
type RAM [ivm.RAMNumFrames]ivm.Frame

// MakeRAM makes an initial RAM module for the virtual machine.
func MakeRAM() RAM {
	r := RAM{}
	for fnum := range r {
		r[fnum] = ivm.MakeFrame()
	}
	return r
}

// AddressFetchWord returns the word value at the given address.
func (r RAM) AddressFetchWord(addr ivm.Address) ivm.Word {
	framenum, frameaddr := ivm.FrameComponentsForAddress(addr / 4)
	return r[framenum][frameaddr]
}

// AddressWriteWord writes the given word value to the given address.
func (r *RAM) AddressWriteWord(addr ivm.Address, val ivm.Word) {
	framenum, frameaddr := ivm.FrameComponentsForAddress(addr / 4)
	r[framenum][frameaddr] = val
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

// FrameFetch fetches the frame with the given frame number.
func (r RAM) FrameFetch(frameNum ivm.FrameNumber) ivm.Frame {
	return r[frameNum]
}

// FrameWrite writes the frame at the given frame number.
func (r *RAM) FrameWrite(frameNum ivm.FrameNumber, frame ivm.Frame) {
	copy(r[frameNum][:], frame[:])
}

// Print prints the contents of RAM to Stdout
func (r RAM) Print() error {
	return r.Fprint(os.Stdout)
}

// Fprint prints the contents of RAM to the given writer
func (r RAM) Fprint(w io.Writer) error {
	for fnum, frame := range r {
		fmt.Fprintf(w, "[%02X: ", fnum)
		frame.Fprint(w)
		fmt.Fprint(w, "]")
		if fnum % 2 == 1 {
			fmt.Fprint(w, "\n")
		} else {
			fmt.Fprint(w, "  ")
		}
	}
	return nil
}

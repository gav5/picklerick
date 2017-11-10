package vm

import (
	"fmt"
	"io"
	"os"
	"strings"

	"./ivm"
)

// Disk describes the virtual machine's disk module.
type Disk [ivm.DiskNumFrames]ivm.Frame

// MakeDisk makes an initial Disk module for the virtual machine.
func MakeDisk() Disk {
	d := Disk{}
	for fnum := range d {
		d[fnum] = ivm.MakeFrame()
	}
	return d
}

// AddressFetchWord returns the word value at the given address.
func (d Disk) AddressFetchWord(addr ivm.Address) ivm.Word {
	framenum, frameaddr := ivm.FrameComponentsForAddress(addr / 4)
	return d[framenum][frameaddr]
}

// AddressWriteWord writes the given word value to the given address.
func (d *Disk) AddressWriteWord(addr ivm.Address, val ivm.Word) {
	framenum, frameaddr := ivm.FrameComponentsForAddress(addr / 4)
	d[framenum][frameaddr] = val
}

// AddressFetchUint32 returns the uint32 value at the given address.
func (d Disk) AddressFetchUint32(addr ivm.Address) uint32 {
	return d.AddressFetchWord(addr).Uint32()
}

// AddressWriteUint32 writes the given uint32 value to the given address.
func (d *Disk) AddressWriteUint32(addr ivm.Address, val uint32) {
	d.AddressWriteWord(addr, ivm.WordFromUint32(val))
}

// AddressFetchInt32 returns the int32 value at the given address.
func (d Disk) AddressFetchInt32(addr ivm.Address) int32 {
	return d.AddressFetchWord(addr).Int32()
}

// AddressWriteInt32 writes the given int32 value to the given address.
func (d *Disk) AddressWriteInt32(addr ivm.Address, val int32) {
	d.AddressWriteWord(addr, ivm.WordFromInt32(val))
}

// AddressFetchBool returns the bool value at the given address.
func (d Disk) AddressFetchBool(addr ivm.Address) bool {
	return d.AddressFetchWord(addr).Bool()
}

// AddressWriteBool writes the given bool value to the given address.
func (d *Disk) AddressWriteBool(addr ivm.Address, val bool) {
	d.AddressWriteWord(addr, ivm.WordFromBool(val))
}

// FrameFetch fetches the frame with the given frame number.
func (d Disk) FrameFetch(frameNum ivm.FrameNumber) ivm.Frame {
	return d[frameNum]
}

// FrameWrite writes the frame at the given frame number.
func (d *Disk) FrameWrite(frameNum ivm.FrameNumber, frame ivm.Frame) {
	copy(d[frameNum][:], frame[:])
}

// Print prints the contents of RAM to Stdout
func (d Disk) Print() error {
	return d.Fprint(os.Stdout)
}

// Fprint prints the contents of RAM to the given writer
func (d Disk) Fprint(w io.Writer) error {
	border := strings.Repeat("-", 11*ivm.FrameDisplayColumns+3)
	var err error
	fmt.Fprint(w, "\n")
	for fnum, frame := range d {
		if fnum > 0 {
			if _, err = fmt.Fprintln(w, border); err != nil {
				return err
			}
		}
		if _, err = fmt.Fprintf(w, "Frame %02X\n", fnum); err != nil {
			return err
		}
		frame.Fprint(w)
	}
	return nil
}

package vm

import (
	"fmt"
	"io"
	"os"

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
	for fnum, frame := range d {
		fmt.Fprintf(w, "[%03X: ", fnum)
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

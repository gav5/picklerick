package vm

import (
	"fmt"
	"io"
	"log"
	"os"

	"../util/logger"
	"./ivm"
)

// Disk describes the virtual machine's disk module.
// type Disk [ivm.DiskNumFrames]ivm.Frame
type Disk struct {
	contents [ivm.DiskNumFrames]ivm.Frame
	logger   *log.Logger
}

// MakeDisk makes an initial Disk module for the virtual machine.
func MakeDisk() Disk {
	d := Disk{
		contents: [ivm.DiskNumFrames]ivm.Frame{},
		logger:   logger.New("disk"),
	}
	for fnum := range d.contents {
		d.contents[fnum] = ivm.MakeFrame()
	}
	return d
}

// MockDisk makes a fake disk for testing.
func MockDisk() Disk {
	d := Disk{
		contents: [ivm.DiskNumFrames]ivm.Frame{},
		logger:   logger.Dummy(),
	}
	for fnum := range d.contents {
		d.contents[fnum] = ivm.MakeFrame()
	}
	return d
}

// FrameFetch fetches the frame with the given frame number.
func (d Disk) FrameFetch(frameNum ivm.FrameNumber) ivm.Frame {
	return d.contents[frameNum]
}

// FrameWrite writes the frame at the given frame number.
func (d *Disk) FrameWrite(frameNum ivm.FrameNumber, frame ivm.Frame) {
	d.logger.Printf("write to frame %03X: %v", int(frameNum), frame)
	copy(d.contents[frameNum][:], frame[:])
}

// Print prints the contents of RAM to Stdout
func (d Disk) Print() error {
	return d.Fprint(os.Stdout)
}

// Fprint prints the contents of RAM to the given writer
func (d Disk) Fprint(w io.Writer) error {
	for fnum, frame := range d.contents {
		var err error

		_, err = fmt.Fprintf(w, "\n[%03X: ", fnum)
		if err != nil {
			return err
		}
		err = frame.Fprint(w)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(w, "]")
		if err != nil {
			return err
		}
	}
	return nil
}

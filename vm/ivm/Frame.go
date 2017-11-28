package ivm

import (
	"fmt"
	"io"
	"math"
)

const (
	// FrameSize is the size of a frame.
	FrameSize = 4

	// FrameDisplayColumns is the number of columns to display a frame with.
	FrameDisplayColumns = 4
)

// Frame describes a frame in the virtual machine.
// (these can apply either to RAM or disk and are interchangeable)
type Frame [FrameSize]Word

// MakeFrame makes an initial frame for the virtual machine.
func MakeFrame() Frame {
	return Frame{
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
	}
}

// Copy makes a duplicate frame.
func (f Frame) Copy() Frame {
	c := Frame{}
	copy(c[:], f[:])
	return c
}

// FrameArrayFromWordArray builds a frame array for a given array of words.
func FrameArrayFromWordArray(ary []Word) []Frame {
	numFrames := int(math.Ceil(float64(len(ary)) / FrameSize))
	frameArray := make([]Frame, numFrames)
	for i, w := range ary {
		frameNum := i / FrameSize
		frameAddr := i % FrameSize
		frameArray[frameNum][frameAddr] = w
	}
	return frameArray
}

// FrameArrayFromUint32Array builds a frame array for a given array of uint32's.
func FrameArrayFromUint32Array(ary []uint32) []Frame {
	numFrames := int(math.Ceil(float64(len(ary)) / FrameSize))
	frameArray := make([]Frame, numFrames)
	for i, w := range ary {
		frameNum := i / FrameSize
		frameAddr := i % FrameSize
		frameArray[frameNum][frameAddr] = WordFromUint32(w)
	}
	return frameArray
}

// Fprint prints the contents of the RAMFrame to the given writer.
func (frame Frame) Fprint(w io.Writer) error {
	var err error
	for faddr, val := range frame {
		switch faddr % FrameDisplayColumns {
		case 0:
			if _, err = fmt.Fprintf(w, "%v", val); err != nil {
				return err
			}
		default:
			if _, err = fmt.Fprintf(w, " | %v", val); err != nil {
				return err
			}
		}
	}
	return nil
}

// FrameComponentsForAddress returns the frame number and frame address for the address.
func FrameComponentsForAddress(addr Address) (FrameNumber, FrameAddress) {
	return FrameNumber(addr / FrameSize), FrameAddress(addr % FrameSize)
}

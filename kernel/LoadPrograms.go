package kernel

import (
	"../prog"
	"../vm/ivm"
)

// LoadPrograms loads the given programs into the virtual machine.
func LoadPrograms(vm ivm.IVM, programs []prog.Program) error {
	cur := ivm.FrameNumber(0)
	for _, p := range programs {
		frames, err := p.Frames()
		if err != nil {
			return err
		}
		if (int(cur) + len(frames)) < ivm.RAMNumFrames {
			for _, f := range frames {
				vm.RAMFrameWrite(cur, f)
				cur++
			}
		} else if (cur - ivm.RAMNumFrames) < ivm.DiskNumFrames {
			for _, f := range frames {
				vm.DiskFrameWrite(cur-ivm.RAMNumFrames, f)
				cur++
			}
		} else {
			return ProgramOverflowError{}
		}
	}
	return nil
}

// ProgramOverflowError means there isn't enough storage to hold all provided programs.
type ProgramOverflowError struct{}

func (e ProgramOverflowError) Error() string {
	return "There isn't enough storage to hold all the provided programs."
}

package ivm

const (
	// DiskNumWords is the number of words in the virtual machine disk.
	DiskNumWords = 2048

	// DiskNumFrames is the number of frames in the virtual machine disk.
	DiskNumFrames = DiskNumWords / FrameSize
)

// IDisk is the kernel interface into a virtual machine disk storage drive.
// (this is because vm uses the kernel, so we have to avoid circular dependencies)
// (this also helps regulate the access of the vm by the kernel to avoid abuse)
type IDisk interface {
	// Frame management
	FrameFetch(FrameNumber) Frame
	FrameWrite(FrameNumber, Frame)
}

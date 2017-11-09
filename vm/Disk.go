package vm

// DiskFrames is the number of frames in the virtual machine disk.
const DiskFrames = 2048

// Disk describes the virtual machine's disk module.
type Disk [DiskFrames]uint32

package cpu

// ID represents the ID of a given CPU
// (this is used to identify a given CPU)
type ID uint8

const (
	// CPU1 represents the first CPU in the system
	CPU1 ID = iota
	// CPU2 represents the second CPU in the system
	CPU2
	// CPU3 represents the third CPU in the system
	CPU3
	// CPU4 represents the fourth CPU in the system
	CPU4
)

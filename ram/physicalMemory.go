package ram

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var physicalMemory = make([]uint32, MemorySize)

// PrintPhysicalMemory prints the contents of the physical memory to Stdout
func PrintPhysicalMemory() error {
	return FprintPhysicalMemory(os.Stdout)
}

// FprintPhysicalMemory prints the contents of the physical memory to the given writer
func FprintPhysicalMemory(w io.Writer) error {
	const numcolumns = 8
	border := strings.Repeat("-", 11*numcolumns+3)
	var err error
	fmt.Fprint(w, "\n")
	for index, val := range physicalMemory {
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

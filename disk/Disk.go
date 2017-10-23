package disk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"

	"../prog"
)

const (
	progTableSize = 32

	// FrameSize defines the number of words in a frame
	FrameSize = 32

	// DiskSize determines how many words are on disk
	DiskSize = 2048
)

var (
	disk = make([]uint32, DiskSize)

	// Order describes the binary order used to encode binary data to disk
	Order = binary.BigEndian
)

type (
	programsTableEntry struct {
		jobID  uint8
		start  uint16
		offset uint8
	}
	tableFullError          struct{}
	entryNotFoundError      struct{}
	entryAlreadyExistsError struct{}
)

func (e programsTableEntry) end() uint16 {
	return e.start + uint16(e.offset) + 1
}

func (e tableFullError) Error() string {
	return "Programs entry table is full"
}
func (e entryNotFoundError) Error() string {
	return "Job ID not found in the programs entry table"
}
func (e entryAlreadyExistsError) Error() string {
	return "This job entry already exists (cannot overwrite the job entry table)"
}

func getProgramTableSize() uint8 {
	var pgTableSize uint8
	for pgTableSize = 0; pgTableSize < progTableSize; pgTableSize++ {
		if disk[pgTableSize] == 0x00000000 {
			break
		}
	}
	return pgTableSize
}

func getProgramsTable() ([]programsTableEntry, error) {
	size := getProgramTableSize()
	fmt.Printf("size: %d\n", size)
	t := make([]programsTableEntry, size)
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, Order, disk[:size]); err != nil {
		return []programsTableEntry{}, err
	}
	fmt.Printf("buf: %x\n", buf)
	if err := binary.Read(buf, Order, &t); err != nil {
		return []programsTableEntry{}, err
	}
	fmt.Printf("wrote to table len(t): %v\n", len(t))
	return t, nil
}

func setProgramsTable(t []programsTableEntry) error {
	buf := new(bytes.Buffer)
	outary := make([]uint32, len(t))
	if err := binary.Write(buf, Order, t); err != nil {
		return err
	}
	if err := binary.Read(buf, Order, outary); err != nil {
		return err
	}
	copy(disk, outary)
	return nil
}

func findProgEntry(jobID uint8, t []programsTableEntry) (programsTableEntry, error) {
	var entry programsTableEntry
	fmt.Printf("ptable: %v\n", t)
	for _, entry = range t {
		if entry.jobID == jobID {
			return entry, nil
		}
	}
	return programsTableEntry{}, entryNotFoundError{}
}

func assignProgEntry(t *[]programsTableEntry, e programsTableEntry) error {
	// make sure it does not exist (returns an error)
	// otherwise, this already exists and we have a problem
	if _, err := findProgEntry(e.jobID, *t); err == nil {
		return entryAlreadyExistsError{}
	}
	// make sure the table isn't already full
	// (because it has a finite length)
	if len(*t) >= progTableSize {
		return tableFullError{}
	}
	// append to the entries table that was passed in
	// (since this is passed by reference, it should affect that value)
	(*t) = append(*t, e)
	return nil
}

// LoadProgram loads the program with the corresponding job ID from disk
func LoadProgram(p *prog.Program, jobID uint8) error {
	var (
		programsTable []programsTableEntry
		// progEntry     programsTableEntry
		// words         []uint32
		err error
	)
	if programsTable, err = getProgramsTable(); err != nil {
		return err
	}
	fmt.Printf("ptable: %v\n", programsTable)
	// if progEntry, err = findProgEntry(jobID, programsTable); err != nil {
	// 	return err
	// }
	// copy(words, disk[progEntry.start:progEntry.end()])
	// fmt.Printf("words: %v\n", words)
	// p.SetWords(words)
	return nil
}

// StoreProgram stores the given progam on the disk
func StoreProgram(program prog.Program) error {
	var (
		progTable []programsTableEntry
		words     []uint32
		newStart  uint16
		err       error
	)
	if progTable, err = getProgramsTable(); err != nil {
		return err
	}
	if words, err = program.GetWords(); err != nil {
		return err
	}
	if len(progTable) > 0 {
		lastEntry := progTable[len(progTable)-1]
		newStart = lastEntry.start + uint16(lastEntry.offset)
	} else {
		newStart = progTableSize
	}
	entry := programsTableEntry{
		jobID:  program.Job.ID,
		start:  newStart,
		offset: uint8(len(words)),
	}
	if err = assignProgEntry(&progTable, entry); err != nil {
		return err
	}
	if err = setProgramsTable(progTable); err != nil {
		return err
	}
	copy(disk[entry.start:entry.end()], words[:])
	return nil
}

// PrintPhysicalDisk prints the contents of the physical disk to Stdout
func PrintPhysicalDisk() error {
	return FprintPhysicalDisk(os.Stdout)
}

// FprintPhysicalDisk prints the contents of the physical disk to the indicated writer
func FprintPhysicalDisk(w io.Writer) error {
	const numcolumns = 8
	border := strings.Repeat("-", 11*numcolumns+3)
	var err error
	fmt.Fprint(w, "\n")
	for index, val := range disk {
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

package prog

import (
	"fmt"
	"os"
)

// ASMFile encapsulates an assembly file
type ASMFile struct {
	file *os.File
}

// MakeASMFile build an assembly file for the given filename
func MakeASMFile(name string) (ASMFile, error) {
	var file *os.File
	var err error
	if file, err = os.Create(name); err != nil {
		return ASMFile{}, err
	}
	return ASMFile{file: file}, nil
}

// WritePrograms writes the given programs array to the appropriate file
func (f ASMFile) WritePrograms(pa []Program, source string) error {
	fmt.Fprintf(f.file, "Extracted %d programs from \"%s\"\n", len(pa), source)
	fmt.Fprintln(f.file, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for _, p := range pa {
		if err := p.WriteASM(f.file); err != nil {
			return err
		}
		fmt.Fprintln(f.file, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
	return nil
}

// Close closes the wrapped file
func (f ASMFile) Close() error {
	return f.file.Close()
}

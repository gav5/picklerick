package prog

import (
	"fmt"
	"io"

	"../instrDecode"
)

// WriteASM writes the assembly instructions to the given file writer
func (p Program) WriteASM(w io.Writer) error {
	p.Job.WriteASM(w)
	return nil
}

// WriteASM writes the assembly instructions to the given file writer
func (j Job) WriteASM(w io.Writer) error {
	fmt.Fprintf(w, "Job ID: %d\n", j.ID)
	for index, iraw := range j.Instructions {
		instr, err := instrDecode.FromUint32(iraw)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%08X  | %04X |  %s\n", iraw, (index * 4), instr.ASM())
	}
	return nil
}

package decoder

import (
	"fmt"
	"io"

	"../prog"
)

// ProgramWriteASM writes the assembly instructions to the given file writer
func ProgramWriteASM(w io.Writer, p prog.Program) error {
	JobWriteASM(w, p.Job)
	return nil
}

// JobWriteASM writes the assembly instructions to the given file writer
func JobWriteASM(w io.Writer, j prog.Job) error {
	fmt.Fprintf(w, "Job ID: %d\n", j.ID)
	for index, iraw := range j.Instructions {
		instr, err := InstrFromUint32(iraw)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%08X  | %04X |  %s\n", iraw, (index * 4), instr.ASM())
	}
	return nil
}

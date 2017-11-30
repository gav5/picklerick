package report

import (
	"fmt"
	"io"

	"../vm"
	"../vm/core"
)

type coreDumpReport struct {
	snapshot core.Snapshot
	coreNum  uint8
	index    int
}

func makeCoreDumpReports(v *vm.VM) ([]reportBase, reportBuilder, error) {
	reports := []reportBase{}
	for _, c := range v.Cores {
		for i, s := range c.Snapshots() {
			reports = append(reports, coreDumpReport{s, c.CoreNum, i + 1})
		}
	}
	return reports, txtReportBuilder{}, nil
}

func (coreDumpReport) Namespace() string {
	return "snapshots"
}

func (r coreDumpReport) Name() string {
	return fmt.Sprintf(
		"snap%02d-%04x.c%d-%d.dump",
		r.snapshot.Process.Program.JobID,
		uint32(r.snapshot.Process.State().ProgramCounter),
		r.coreNum, r.index,
	)
}

func (r coreDumpReport) Title() string {
	return fmt.Sprintf(
		"Virtual Machine Core Dump (%02d@%04x; Core %d, Snapshot %d)",
		r.snapshot.Process.Program.JobID,
		uint32(r.snapshot.Process.State().ProgramCounter),
		r.coreNum, r.index,
	)
}

func (r coreDumpReport) Fprint(w io.Writer) error {
	var err error

	proc := r.snapshot.Process
	prog := proc.Program
	currentState := proc.State()
	nextState := r.snapshot.Next

	err = fprintProperty(w, "Process Number", proc.ProcessNumber)
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Priority", proc.Priority)
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Program Counter", currentState.ProgramCounter)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Summarized Result")
	if err != nil {
		return err
	}
	if nextState.Error != nil {
		_, err = fmt.Fprintf(
			w, "\nERROR: %v",
			nextState.Error,
		)
		if err != nil {
			return err
		}
	} else if nextState.Halt {
		_, err = fmt.Fprint(w, "\nHALT")
		if err != nil {
			return err
		}
	} else if len(nextState.Faults) > 0 {
		_, err = fmt.Fprintf(
			w, "\nFAULTS at %v",
			nextState.Faults,
		)
		if err != nil {
			return err
		}
	} else if nextState.ProgramCounter != (currentState.ProgramCounter + 4) {
		_, err = fmt.Fprintf(
			w, "\nJUMP to 0x%04X",
			uint32(nextState.ProgramCounter),
		)
		if err != nil {
			return err
		}
	} else {
		_, err = fmt.Fprintf(
			w, "\nCONTINUE to 0x%04X",
			uint32(nextState.ProgramCounter),
		)
		if err != nil {
			return err
		}
	}

	err = fprintHeader(w, "Decoded Assembly")
	if err != nil {
		return err
	}
	err = prog.FprintInstructionsPC(w, currentState.ProgramCounter)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Registers Before")
	if err != nil {
		return err
	}
	err = currentState.FprintRegisters(w)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Caches Before")
	if err != nil {
		return err
	}
	err = currentState.Caches.Fprint(w)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Registers After")
	if err != nil {
		return err
	}
	err = nextState.FprintRegisters(w)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Caches After")
	if err != nil {
		return err
	}
	err = nextState.Caches.Fprint(w)
	if err != nil {
		return err
	}

	return nil
}

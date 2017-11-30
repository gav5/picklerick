package report

import (
	"fmt"
	"io"

	"../vm"
)

type ramDumpReport struct {
	ram vm.RAM
}

func makeRAMDumpReport(v *vm.VM) ([]reportBase, reportBuilder, error) {
	return []reportBase{ramDumpReport{v.RAM}}, txtReportBuilder{}, nil
}

func (ramDumpReport) Namespace() string {
	return "vm"
}

func (ramDumpReport) Name() string {
	return "ram.dump"
}

func (ramDumpReport) Title() string {
	return "Virtual Machine RAM Dump"
}

func (r ramDumpReport) Fprint(w io.Writer) error {
	for fnum, frame := range r.ram {
		var err error

		_, err = fmt.Fprintf(w, "\n[%02X: ", fnum)
		if err != nil {
			return err
		}
		err = frame.Fprint(w)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(w, "]")
		if err != nil {
			return err
		}
	}
	return nil
}

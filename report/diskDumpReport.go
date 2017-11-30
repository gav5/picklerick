package report

import (
	"fmt"
	"io"

	"../vm"
)

type diskDumpReport struct {
	disk vm.Disk
}

func makeDiskDumpReport(v *vm.VM) ([]reportBase, reportBuilder, error) {
	return []reportBase{diskDumpReport{v.Disk}}, txtReportBuilder{}, nil
}

func (diskDumpReport) Namespace() string {
	return "vm"
}

func (diskDumpReport) Name() string {
	return "disk.dump"
}

func (diskDumpReport) Title() string {
	return "Virtual Machine Disk Dump"
}

func (r diskDumpReport) Fprint(w io.Writer) error {
	for fnum, frame := range r.disk {
		var err error

		_, err = fmt.Fprintf(w, "\n[%03X: ", fnum)
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

package report

import (
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
	return r.disk.Fprint(w)
}

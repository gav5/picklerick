package report

import (
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
	return r.ram.Fprint(w)
}

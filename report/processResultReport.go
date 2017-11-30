package report

import (
	"fmt"
	"io"

	"../kernel/process"
	"../vm"
	"../vm/decoder"
)

type processResultReport struct {
	p process.Process
}

func makeProcessResultReports(v *vm.VM) ([]reportBase, reportBuilder, error) {
	processTable := v.ProcessTable()
	pTableLen := len(processTable)
	reports := make([]reportBase, pTableLen)
	for i, p := range processTable {
		reports[pTableLen-i-1] = processResultReport{p}
	}
	return reports, txtReportBuilder{}, nil
}

func (processResultReport) Namespace() string {
	return "process"
}

func (r processResultReport) Name() string {
	return fmt.Sprintf("process%d.result", r.p.ProcessNumber)
}

func (r processResultReport) Title() string {
	return fmt.Sprintf("Process %d Result Report", r.p.ProcessNumber)
}

func (r processResultReport) Fprint(w io.Writer) error {

	fprintProperty(w, "CodeSize", r.p.ProcessNumber)
	fprintProperty(w, "Priority", r.p.Priority)

	fprintHeader(w, "Decoded Assembly")
	for i, iRAW := range r.p.Program.Instructions {
		iDecoded, err := decoder.DecodeInstruction(iRAW)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\n%04X | %08X | %s", i*4, iRAW, iDecoded.Assembly())
	}

	return nil
}

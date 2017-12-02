package report

import (
	"../kernel/process"
	"../vm"
)

type processTableReport struct {
	pt []process.Process
}

func makeProcessTableReport(v *vm.VM) ([]reportBase, reportBuilder, error) {
	return []reportBase{
		processTableReport{pt: v.ProcessTable()},
	}, csvReportBuilder{}, nil
}

func (processTableReport) Namespace() string {
	return ""
}

func (processTableReport) Name() string {
	return "processTable"
}

func (r processTableReport) TableHeaders() []string {
	return process.TableHeaders()
}

func (r processTableReport) TableRows() [][]string {
	ptLen := len(r.pt)
	dataRows := make([][]string, ptLen)
	for rowNum, p := range r.pt {
		// fill in reverse (because the table is in reverse order)
		dataRows[ptLen-rowNum-1] = p.TableRow()
	}
	return dataRows
}

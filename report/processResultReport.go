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

	inBufStart, inBufEnd := r.p.InputBufferRange()
	fprintHeader(w, fmt.Sprintf(
		"Input Buffer [0x%02X - 0x%02X]", uint32(inBufStart), uint32(inBufEnd),
	))
	fprintWords(w, r.p.InputBuffer())

	tempBufStart, tempBufEnd := r.p.TempBufferRange()
	fprintHeader(w, fmt.Sprintf(
		"Temp Buffer [0x%02X - 0x%02X]", uint32(tempBufStart), uint32(tempBufEnd),
	))
	fprintWords(w, r.p.TempBuffer())

	outBufStart, outBufEnd := r.p.OutputBufferRange()
	fprintHeader(w, fmt.Sprintf(
		"Output Buffer [0x%02X - 0x%02X]", uint32(outBufStart), uint32(outBufEnd),
	))
	fprintWords(w, r.p.OutputBuffer())

	return nil
}

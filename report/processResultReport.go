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
	v *vm.VM
}

func makeProcessResultReports(v *vm.VM) ([]reportBase, reportBuilder, error) {
	processTable := v.ProcessTable()
	pTableLen := len(processTable)
	reports := make([]reportBase, pTableLen)
	for i, p := range processTable {
		reports[pTableLen-i-1] = processResultReport{p, v}
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
	var err error

	err = fprintProperty(w, "CodeSize", r.p.ProcessNumber)
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Priority", r.p.Priority)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Decoded Assembly")
	if err != nil {
		return err
	}
	for i, iRAW := range r.p.Program.Instructions {
		iDecoded, err := decoder.DecodeInstruction(iRAW)
		if err != nil {
			return err
		}
		asm := iDecoded.Assembly()
		_, err = fmt.Fprintf(w, "\n%04X | %08X | %s", i*4, iRAW, asm)
		if err != nil {
			return err
		}
	}

	inBufStart, inBufEnd := r.p.InputBufferRange()
	err = fprintHeader(w, fmt.Sprintf(
		"Input Buffer [0x%02X - 0x%02X]", uint32(inBufStart), uint32(inBufEnd),
	))
	if err != nil {
		return err
	}
	err = fprintWords(w, r.p.InputBuffer(r.v))
	if err != nil {
		return err
	}

	tempBufStart, tempBufEnd := r.p.TempBufferRange()
	err = fprintHeader(w, fmt.Sprintf(
		"Temp Buffer [0x%02X - 0x%02X]", uint32(tempBufStart), uint32(tempBufEnd),
	))
	if err != nil {
		return err
	}
	err = fprintWords(w, r.p.TempBuffer(r.v))
	if err != nil {
		return err
	}

	outBufStart, outBufEnd := r.p.OutputBufferRange()
	err = fprintHeader(w, fmt.Sprintf(
		"Output Buffer [0x%02X - 0x%02X]", uint32(outBufStart), uint32(outBufEnd),
	))
	if err != nil {
		return err
	}
	err = fprintWords(w, r.p.OutputBuffer(r.v))
	if err != nil {
		return err
	}

	err = fprintHeader(w, "RAM Page Table")
	if err != nil {
		return err
	}
	err = r.p.RAMPageTable.Fprint(w)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Disk Page Table")
	if err != nil {
		return err
	}
	err = r.p.DiskPageTable.Fprint(w)
	if err != nil {
		return err
	}

	return nil
}

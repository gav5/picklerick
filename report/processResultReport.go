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

	err = fprintProperty(w, "Priority", r.p.Priority)
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Status", r.p.Status())
	if err != nil {
		return err
	}
	if r.p.Status() == process.Terminated {
		// Completion only listed if it was completed!
		err = fprintProperty(w, "Completion Time", r.p.Metrics.CompletionTime.Value())
		if err != nil {
			return err
		}
	}
	err = fprintProperty(w, "Wait Time", r.p.Metrics.WaitTime.Value())
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Cycles Used", r.p.Metrics.Cycles.Value())
	if err != nil {
		return err
	}
	err = fprintProperty(w, "Max Cache Size", r.p.Metrics.CacheSize.Max())
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

	err = fprintHeader(w, "Input Buffer")
	if err != nil {
		return err
	}
	inWords, inOffset := r.p.InputBuffer(r.v)
	err = fprintBuffer(w, inWords, inOffset)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Output Buffer")
	if err != nil {
		return err
	}
	outWords, outOffset := r.p.OutputBuffer(r.v)
	err = fprintBuffer(w, outWords, outOffset)
	if err != nil {
		return err
	}

	err = fprintHeader(w, "Temp Buffer")
	if err != nil {
		return err
	}
	tempWords, tempOffset := r.p.TempBuffer(r.v)
	err = fprintBuffer(w, tempWords, tempOffset)
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

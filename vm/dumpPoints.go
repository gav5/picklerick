package vm

import (
	"fmt"
	"io"
	"strings"

	"../kernel/process"
	"./ivm"
)

type dumpPoint struct {
	jobID          uint8
	programCounter ivm.Address
}

type dumpPointsList []dumpPoint

func (dpl dumpPointsList) ShouldDump(p process.Process) bool {
	jid, pc := p.Program.JobID, p.State().ProgramCounter
	for _, dp := range dpl {
		if dp.jobID == jid && dp.programCounter == pc {
			return true
		}
	}
	return false
}

func makeDumpPoints(svar string) (dumpPointsList, error) {
	rawpoints := strings.Split(svar, ",")
	retary := dumpPointsList{}
	for _, rp := range rawpoints {
		dp := dumpPoint{}
		_, err := fmt.Sscanf(rp, "%d@%x", &dp.jobID, &dp.programCounter)
		if err == io.EOF {
			continue // just ignore these
		}
		if err != nil {
			return retary, err
		}
		retary = append(retary, dp)
	}
	return retary, nil
}

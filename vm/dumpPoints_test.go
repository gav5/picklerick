package vm

import (
	"reflect"
	"testing"

	"../kernel/process"
	"./ivm"
)

var makeDumpPointsTests = []struct {
	in  string
	out dumpPointsList
	err bool
}{
	{in: "", out: []dumpPoint{}, err: false},
	{in: "2@4C", out: []dumpPoint{{2, 0x4c}}, err: false},
	{in: "1@04,2@4a", out: []dumpPoint{{1, 0x04}, {2, 0x4a}}, err: false},
	{in: "4c", out: []dumpPoint{}, err: true},
	{in: "1@4c,", out: []dumpPoint{{1, 0x4c}}, err: false},
}

func TestMakeDumpPoints(t *testing.T) {
	for _, tt := range makeDumpPointsTests {
		out, err := makeDumpPoints(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf(
				"makeDumpPoints(\"%s\") => %v,_ (should be %v)",
				tt.in, out, tt.out,
			)
		}
		if tt.err && (err == nil) {
			t.Errorf(
				"makeDumpPoints(\"%s\") => _, %v (should be a failure)",
				tt.in, err,
			)
		} else if !tt.err && (err != nil) {
			t.Error(
				"makeDumpPoints(\"%s\") => _, %v (should not fail here)",
				tt.in, err,
			)
		}
	}
}

var shouldDumpTests = []struct {
	dpl   dumpPointsList
	jobID uint8
	pc    ivm.Address
	out   bool
}{
	{
		dpl: dumpPointsList{
			{1, 0x3e}, {1, 0x42}, {4, 0x55},
		},
		jobID: 1,
		pc:    0x42,
		out:   true,
	},
	{
		dpl: dumpPointsList{
			{1, 0x3e}, {1, 0x42}, {4, 0x55},
		},
		jobID: 1,
		pc:    0xff,
		out:   false,
	},
	{
		dpl: dumpPointsList{
			{1, 0x3e}, {1, 0x42}, {4, 0x55},
		},
		jobID: 2,
		pc:    0x42,
		out:   false,
	},
}

func TestShouldDump(t *testing.T) {
	for _, tt := range shouldDumpTests {
		p := process.Sleep()
		p.Program.JobID = tt.jobID
		p.SetState(ivm.State{ProgramCounter: tt.pc})
		out := tt.dpl.ShouldDump(p)
		if out != tt.out {
			t.Errorf(
				"[%v].ShouldDump(%d@%x) => %v (should be %v)",
				tt.dpl, tt.jobID, tt.pc, out, tt.out,
			)
		}
	}
}

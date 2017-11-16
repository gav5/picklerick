package prog

// import (
// 	"reflect"
// 	"testing"
// )
//
// var (
// 	p1 = Program{Job: Job{ID: 0x1}}
// 	p2 = Program{Job: Job{ID: 0x2}}
// 	p3 = Program{Job: Job{ID: 0x3}}
// 	p4 = Program{Job: Job{ID: 0x4}}
// )
//
// var makeScheduleTests = []struct {
// 	programsIn  []Program
// 	programsOut []Program
// }{
// 	{
// 		// An empty list should yield and empty list
// 		[]Program{},
// 		[]Program{},
// 	},
// 	{
// 		// A list of four jobs should yield the same list
// 		[]Program{p1, p2, p3, p4},
// 		[]Program{p1, p2, p3, p4},
// 	},
// }
//
// func TestMakeSchedule(t *testing.T) {
// 	for _, tt := range makeScheduleTests {
// 		sc := MakeSchedule(tt.programsIn)
// 		if !reflect.DeepEqual(sc.programsList, tt.programsOut) {
// 			t.Errorf("MakeSchedule(%v) => %v; want %v", tt.programsIn, sc.programsList, tt.programsOut)
// 		}
// 	}
// }
//
// func TestCurrentProgram(t *testing.T) {
// 	sc := MakeSchedule([]Program{p1, p2, p3, p4})
// 	currentProgram := *sc.CurrentProgram()
// 	if !reflect.DeepEqual(currentProgram, p1) {
// 		t.Errorf("(Schedule) CurrentProgram() => %v; want %v", currentProgram, p1)
// 	}
// }
//
// var nextProgramTests = []struct {
// 	programsIn         []Program
// 	programOutSequence []Program
// }{
// 	{
// 		[]Program{p1, p2, p3, p4},
// 		[]Program{p1, p2, p3, p4},
// 	},
// 	{
// 		[]Program{p4, p2, p3, p1},
// 		[]Program{p4, p2, p3, p1},
// 	},
// }
//
// func TestNextProgram(t *testing.T) {
// 	var currentProgram *Program
// 	for _, tt := range nextProgramTests {
// 		sc := MakeSchedule(tt.programsIn)
// 		currentProgram = sc.CurrentProgram()
// 		for i, p := range tt.programOutSequence {
// 			if i > 0 {
// 				currentProgram = sc.NextProgram()
// 			}
// 			if !reflect.DeepEqual(*currentProgram, p) {
// 				t.Errorf("(Schedule) NextProgram() => %v; want %v", *currentProgram, p)
// 			}
// 		}
// 		currentProgram = sc.NextProgram()
// 		if currentProgram != nil {
// 			t.Errorf("(Schedule) NextProgram() => %v; want %v", currentProgram, nil)
// 		}
// 	}
// }

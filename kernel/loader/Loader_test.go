package loader

import (
	"testing"
)

var loadTests = []struct{
  programFile string
  err error
  numPrograms int
}{
  {
    programFile: "../../Program-File.txt",
    err: nil,
    numPrograms: 30,
  },
}

func TestLoad(t *testing.T) {
  for _, tt := range loadTests {

    // simulate call to Load
    programArray, err := Load(tt.programFile)

    if err != tt.err {
      t.Errorf(
        "[%s].Load() => %v, expected %v\n",
        tt.programFile, err, tt.err,
      )
    }
    numPrograms := len(programArray)
    if numPrograms != tt.numPrograms {
      t.Errorf(
        "[%s].Export => %d programs, expected %d programs\n",
        tt.programFile, numPrograms, tt.numPrograms,
      )
    }
  }
}

// p1, e1 := ParseFile("../Program-File.txt")
// if e1 != nil {
// 	t.Errorf("ParseFile(<real>): expected no error, got %v", e1)
// } else {
// 	if len(p1) != 30 {
// 		t.Errorf("ParseFile(<real>): expected 30 programs, got %d", len(p1))
// 	}
// }
// // test a nonexistant file (should return an error)
// p2, e2 := ParseFile("idontexist.whocares")
// if e2 == nil {
// 	t.Errorf("ParseFile(<wrong>): expected an error, no error was provided")
// }
// if p2 != nil {
// 	t.Errorf("ParseFile(<wrong>): expected nil, got %d progams", len(p2))
// }

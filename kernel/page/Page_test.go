package page

import (
  "reflect"
  "testing"

  "../../vm/ivm"
  "../../prog"
)

var translateAddressTests = []struct {
  pt Table
  in ivm.Address
  out ivm.Address
}{
  {
    pt: Table{0: 0},
    in: 0x00000000,
    out: 0x00000000,
  },
  {
    pt: Table{1: 0},
    in: 0x00000004,
    out: 0x00000000,
  },
  {
    pt: Table{1: 0},
    in: 0x00000005,
    out: 0x00000001,
  },
  {
    pt: Table{1: 0},
    in: 0x00000006,
    out: 0x00000002,
  },
  {
    pt: Table{1: 0},
    in: 0x00000007,
    out: 0x000000003,
  },
}

func TestTranslateAddress(t *testing.T) {
  for _, tt := range translateAddressTests {
    addr := tt.pt.TranslateAddress(tt.in)
    if addr != tt.out {
      t.Errorf(
        "[%v].TranslateAddress(%v) => %v (expected %v)\n",
        tt.pt, tt.in, addr, tt.out,
      )
    }
  }
}

var arrayFromFrameArrayTests = []struct {
  in []ivm.Frame
  out []Page
}{
  {
    in: []ivm.Frame{
      ivm.Frame{0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      ivm.Frame{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0},
      ivm.Frame{0x00000000, 0x00000000, 0x00000000, 0x00000000},
    },
    out: []Page{
      Page{0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
    },
  },
}

func TestArrayFromFrameArray(t *testing.T) {
  for _, tt := range arrayFromFrameArrayTests {
    pageArray := ArrayFromFrameArray(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "ArrayFromFrameArray() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var arrayFromWordArrayTests = []struct {
  in []ivm.Word
  out []Page
}{
  {
    in: []ivm.Word{
      0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
    },
    out: []Page{
      Page{0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0},
    },
  },
}

func TestArrayFromWordArray(t *testing.T) {
  for _, tt := range arrayFromWordArrayTests {
    pageArray := ArrayFromWordArray(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "ArrayFromWordArray() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var arrayFromUint32ArrayTests = []struct {
  in []uint32
  out []Page
}{
  {
    in: []uint32{
      0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
    },
    out: []Page{
      Page{0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF},
      Page{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0},
    },
  },
}

func TestArrayFromUint32Array(t *testing.T) {
  for _, tt := range arrayFromUint32ArrayTests {
    pageArray := ArrayFromUint32Array(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "ArrayFromUint32Array() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var arrayFromProgramTests = []struct {
  program prog.Program
  pageArray []Page
}{
  {
    prog.Program{
  		Job: prog.Job{
  			ID:             0x01,
  			NumberOfWords:  23,
  			PriorityNumber: 0x01,
  			Instructions: []uint32{
          0xC050005C, 0x4B060000, 0x4B010000, 0x4B000000,
        	0x4F0A005C, 0x4F0D00DC, 0x4C0A0004, 0xC0BA0000,
        	0x42BD0000, 0x4C0D0004, 0x4C060001, 0x10658000,
        	0x56810018, 0x4B060000, 0x4F0900DC, 0x43970000,
        	0x05070000, 0x4C060001, 0x4C090004, 0x10658000,
        	0x5681003C, 0xC10000AC, 0x92000000,
  			},
  		},
  		Data: prog.Data{
  			InputBufferSize:  0x05,
  			OutputBufferSize: 0x02,
  			TempBufferSize:   0x03,
  			DataBlock: [44]uint32{
          0x0000000A, 0x00000006, 0x0000002C, 0x00000045,
          0x00000001, 0x00000007, 0x00000000, 0x00000001,
          0x00000005, 0x0000000A, 0x00000055, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
          0x00000000, 0x00000000, 0x00000000, 0x00000000,
  			},
  		},
  	},
    []Page{
      Page{0xC050005C, 0x4B060000, 0x4B010000, 0x4B000000},
      Page{0x4F0A005C, 0x4F0D00DC, 0x4C0A0004, 0xC0BA0000},
      Page{0x42BD0000, 0x4C0D0004, 0x4C060001, 0x10658000},
      Page{0x56810018, 0x4B060000, 0x4F0900DC, 0x43970000},
      Page{0x05070000, 0x4C060001, 0x4C090004, 0x10658000},
      Page{0x5681003C, 0xC10000AC, 0x92000000, 0x0000000A},
      Page{0x00000006, 0x0000002C, 0x00000045, 0x00000001},
      Page{0x00000007, 0x00000000, 0x00000001, 0x00000005},
      Page{0x0000000A, 0x00000055, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000, 0x00000000},
      Page{0x00000000, 0x00000000, 0x00000000},
    },
  },
}

func TestArrayFromProgram(t *testing.T) {
  for _, tt := range arrayFromProgramTests {
    pageArray := ArrayFromProgram(tt.program)
    if !reflect.DeepEqual(pageArray, tt.pageArray) {
      t.Errorf(
        "ArrayFromProgram(pg%d) %s\nhave\t%v\nwant\t%v",
        tt.program.Job.ID, "did not return the expected value",
        pageArray, tt.pageArray,
      )
    }
  }
}

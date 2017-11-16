package kernel

import (
  "reflect"
  "testing"
  "../vm/ivm"

  "../prog"
)

var pageArrayFromFrameArrayTests = []struct {
  in []ivm.Frame
  out []Page
}{
  {
    in: []ivm.Frame{
      ivm.Frame{
        0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      ivm.Frame{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      ivm.Frame{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
    },
    out: []Page{
      Page{
        0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
    },
  },
}

func TestPageArrayFromFrameArray(t *testing.T) {
  for _, tt := range pageArrayFromFrameArrayTests {
    pageArray := PageArrayFromFrameArray(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "PageArrayFromFrameArray() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var pageArrayFromWordArrayTests = []struct {
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
      Page{
        0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
    },
  },
}

func TestPageArrayFromWordArray(t *testing.T) {
  for _, tt := range pageArrayFromWordArrayTests {
    pageArray := PageArrayFromWordArray(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "PageArrayFromWordArray() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var pageArrayFromUint32ArrayTests = []struct {
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
      Page{
        0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
      },
      Page{
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
        0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
    },
  },
}

func TestPageArrayFromUint32Array(t *testing.T) {
  for _, tt := range pageArrayFromUint32ArrayTests {
    pageArray := PageArrayFromUint32Array(tt.in)
    if !reflect.DeepEqual(pageArray, tt.out) {
      t.Errorf(
        "%s\nhave\t%v\nwant\t%v\n",
        "PageArrayFromUint32Array() did not return the expected value",
        pageArray, tt.out,
      )
    }
  }
}

var pageArrayFromProgramTests = []struct {
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
      Page{
        0xC050005C, 0x4B060000, 0x4B010000, 0x4B000000,
        0x4F0A005C, 0x4F0D00DC, 0x4C0A0004, 0xC0BA0000,
        0x42BD0000, 0x4C0D0004, 0x4C060001, 0x10658000,
        0x56810018, 0x4B060000, 0x4F0900DC, 0x43970000,
      },
      Page{
        0x05070000, 0x4C060001, 0x4C090004, 0x10658000,
        0x5681003C, 0xC10000AC, 0x92000000, 0x0000000A,
        0x00000006, 0x0000002C, 0x00000045, 0x00000001,
        0x00000007, 0x00000000, 0x00000001, 0x00000005,
      },
      Page {
        0x0000000A, 0x00000055, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
      Page {
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
        0x00000000, 0x00000000, 0x00000000, 0x00000000,
      },
      Page {
        0x00000000, 0x00000000, 0x00000000,
      },
    },
  },
}

func TestPageArrayFromProgram(t *testing.T) {
  for _, tt := range pageArrayFromProgramTests {
    pageArray := PageArrayFromProgram(tt.program)
    if !reflect.DeepEqual(pageArray, tt.pageArray) {
      t.Errorf(
        "PageArrayFromProgram(pg%d) %s\nhave\t%v\nwant\t%v",
        tt.program.Job.ID, "did not return the expected value",
        pageArray, tt.pageArray,
      )
    }
  }
}

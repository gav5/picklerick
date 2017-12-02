package page

import (
	"reflect"
	"testing"

	"../../vm/ivm"
)

var translateAddressTests = []struct {
	pt  Table
	in  ivm.Address
	out ivm.Address
}{
	{
		pt:  Table{0: 0},
		in:  0x00000000,
		out: 0x00000000,
	},
	{
		pt:  Table{1: 0},
		in:  0x00000004,
		out: 0x00000000,
	},
	{
		pt:  Table{1: 0},
		in:  0x00000005,
		out: 0x00000001,
	},
	{
		pt:  Table{1: 0},
		in:  0x00000006,
		out: 0x00000002,
	},
	{
		pt:  Table{1: 0},
		in:  0x00000007,
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
	in  []ivm.Frame
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
	in  []ivm.Word
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
	in  []uint32
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

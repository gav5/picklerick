package instrDecode

import (
	"testing"

	"../instrType"
)

var fromUint32Tests = []struct {
	in  uint32
	out string
	err error
}{
	{0xC050005C, "RD R5 0x0000005C", nil},
	{0x4B060000, "MOVI R6 0", nil},
	{0x4B010000, "MOVI R1 0", nil},
	{0x4B000000, "MOVI R0 0", nil},
	{0x4F0A005C, "LDI R10 92", nil},
	{0x4F0D00DC, "LDI R13 220", nil},
	{0x4C0A0004, "ADDI R10 4", nil},
	{0xC0BA0000, "RD R11 (R10)", nil},
	{0x42BD0000, "ST (R13) R11", nil},
	{0x4C0D0004, "ADDI R13 4", nil},
	{0x4C060001, "ADDI R6 1", nil},
	{0x10658000, "SLT R8 R6 R5", nil},
	{0x56810018, "BNE R8 R1 0x0018", nil},
	{0x4B060000, "MOVI R6 0", nil},
	{0x4F0900DC, "LDI R9 220", nil},
	{0x43970000, "LW R7 0(R9)", nil},
	{0x05070000, "ADD R0 R0 R7", nil},
	{0x4C060001, "ADDI R6 1", nil},
	{0x4C090004, "ADDI R9 4", nil},
	{0x10658000, "SLT R8 R6 R5", nil},
	{0x5681003C, "BNE R8 R1 0x003C", nil},
	{0xC10000AC, "WR R0 0x000000AC", nil},
	{0x92000000, "HLT", nil},
	{0xFFFFFFFF, "<nil>", UnrecognizedOpcodeError{op: instrType.Opcode(0xff)}},
}

func TestFromUint32(t *testing.T) {
	for _, tt := range fromUint32Tests {
		val, err := FromUint32(tt.in)

		var asm string
		if val != nil {
			asm = val.ASM()
		} else {
			asm = "<nil>"
		}
		if (asm != tt.out) || (err != tt.err) {
			t.Errorf("FromUint32(0x%08X) => %s,%v; want %s,%v", tt.in, asm, err, tt.out, err)
		}
	}
}

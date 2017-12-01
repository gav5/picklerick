package core

import (
	"fmt"
	"testing"

	"../../kernel/process"
	"../../kernel/program"
	"../ivm"
)

// feed states into a core and verify the states it sends back

type regTest map[ivm.RegisterDesignation]ivm.Word
type cacheTest map[ivm.Address]ivm.Word

type testPoint struct {
	pc        ivm.Address
	halt      bool
	err       error
	registers regTest
	caches    cacheTest
}

var coreTests = []struct {
	sampleProgram     program.Program
	stateTestSequence []testPoint
}{
	{
		program.Program{
			JobID:          0x01,
			NumberOfWords:  0x17,
			PriorityNumber: 0x02,
			Instructions: []uint32{
				0xC050005C, 0x4B060000, 0x4B010000, 0x4B000000,
				0x4F0A005C, 0x4F0D00DC, 0x4C0A0004, 0xC0BA0000,
				0x42BD0000, 0x4C0D0004, 0x4C060001, 0x10658000,
				0x56810018, 0x4B060000, 0x4F0900DC, 0x43970000,
				0x05070000, 0x4C060001, 0x4C090004, 0x10658000,
				0x5681003C, 0xC10000AC, 0x92000000,
			},
			InputBufferSize:  0x14,
			OutputBufferSize: 0x0C,
			TempBufferSize:   0x0C,
			DataBlock: []uint32{
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
		[]testPoint{
			/* 0 */ {0x00, false, nil, regTest{5: 0xA}, cacheTest{
				0x5c: 0x0000000A, 0x60: 0x00000006, 0x64: 0x0000002C, 0x68: 0x00000045,
				0x6c: 0x00000001, 0x70: 0x00000007, 0x74: 0x00000000, 0x78: 0x00000001,
				0x7c: 0x00000005, 0x80: 0x0000000A, 0x84: 0x00000055,
			}},
			/*   1 */ {0x04, false, nil, regTest{6: 0}, cacheTest{}},
			/*   2 */ {0x08, false, nil, regTest{1: 0}, cacheTest{}},
			/*   3 */ {0x0c, false, nil, regTest{0: 0}, cacheTest{}},
			/*   4 */ {0x10, false, nil, regTest{10: 0x5c}, cacheTest{}},
			/*   5 */ {0x14, false, nil, regTest{13: 0xdc}, cacheTest{}},
			/*   6 */ {0x18, false, nil, regTest{10: 0x60}, cacheTest{}},
			/*   7 */ {0x1c, false, nil, regTest{11: 0x6}, cacheTest{}},
			/*   8 */ {0x20, false, nil, regTest{}, cacheTest{0xdc: 0x6}},
			/*   9 */ {0x24, false, nil, regTest{13: 0xe0}, cacheTest{}},
			/*  10 */ {0x28, false, nil, regTest{6: 1}, cacheTest{}},
			/*  11 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  12 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  13 */ {0x18, false, nil, regTest{10: 0x64}, cacheTest{}},
			/*  14 */ {0x1c, false, nil, regTest{11: 0x2c}, cacheTest{}},
			/*  15 */ {0x20, false, nil, regTest{}, cacheTest{0xe0: 0x2c}},
			/*  16 */ {0x24, false, nil, regTest{13: 0xe4}, cacheTest{}},
			/*  17 */ {0x28, false, nil, regTest{6: 2}, cacheTest{}},
			/*  18 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  19 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  20 */ {0x18, false, nil, regTest{10: 0x68}, cacheTest{}},
			/*  21 */ {0x1c, false, nil, regTest{11: 0x45}, cacheTest{}},
			/*  22 */ {0x20, false, nil, regTest{}, cacheTest{0xe4: 0x45}},
			/*  23 */ {0x24, false, nil, regTest{13: 0xe8}, cacheTest{}},
			/*  24 */ {0x28, false, nil, regTest{6: 3}, cacheTest{}},
			/*  25 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  26 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  27 */ {0x18, false, nil, regTest{10: 0x6c}, cacheTest{}},
			/*  28 */ {0x1c, false, nil, regTest{11: 0x1}, cacheTest{}},
			/*  29 */ {0x20, false, nil, regTest{}, cacheTest{0xe8: 0x1}},
			/*  30 */ {0x24, false, nil, regTest{13: 0xec}, cacheTest{}},
			/*  31 */ {0x28, false, nil, regTest{6: 4}, cacheTest{}},
			/*  32 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  33 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  34 */ {0x18, false, nil, regTest{10: 0x70}, cacheTest{}},
			/*  35 */ {0x1c, false, nil, regTest{11: 0x7}, cacheTest{}},
			/*  36 */ {0x20, false, nil, regTest{}, cacheTest{0xec: 0x7}},
			/*  37 */ {0x24, false, nil, regTest{13: 0xf0}, cacheTest{}},
			/*  38 */ {0x28, false, nil, regTest{6: 5}, cacheTest{}},
			/*  39 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  40 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  41 */ {0x18, false, nil, regTest{10: 0x74}, cacheTest{}},
			/*  42 */ {0x1c, false, nil, regTest{11: 0x0}, cacheTest{}},
			/*  43 */ {0x20, false, nil, regTest{}, cacheTest{0xf0: 0x0}},
			/*  44 */ {0x24, false, nil, regTest{13: 0xf4}, cacheTest{}},
			/*  45 */ {0x28, false, nil, regTest{6: 6}, cacheTest{}},
			/*  46 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  47 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  48 */ {0x18, false, nil, regTest{10: 0x78}, cacheTest{}},
			/*  49 */ {0x1c, false, nil, regTest{11: 0x1}, cacheTest{}},
			/*  50 */ {0x20, false, nil, regTest{}, cacheTest{0xf4: 0x1}},
			/*  51 */ {0x24, false, nil, regTest{13: 0xf8}, cacheTest{}},
			/*  52 */ {0x28, false, nil, regTest{6: 7}, cacheTest{}},
			/*  53 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  54 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  55 */ {0x18, false, nil, regTest{10: 0x7c}, cacheTest{}},
			/*  56 */ {0x1c, false, nil, regTest{11: 0x5}, cacheTest{}},
			/*  57 */ {0x20, false, nil, regTest{}, cacheTest{0xf8: 0x5}},
			/*  58 */ {0x24, false, nil, regTest{13: 0xfc}, cacheTest{}},
			/*  59 */ {0x28, false, nil, regTest{6: 8}, cacheTest{}},
			/*  60 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  61 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  62 */ {0x18, false, nil, regTest{10: 0x80}, cacheTest{}},
			/*  63 */ {0x1c, false, nil, regTest{11: 0xa}, cacheTest{}},
			/*  64 */ {0x20, false, nil, regTest{}, cacheTest{0xfc: 0xa}},
			/*  65 */ {0x24, false, nil, regTest{13: 0x100}, cacheTest{}},
			/*  66 */ {0x28, false, nil, regTest{6: 9}, cacheTest{}},
			/*  67 */ {0x2c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  68 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  69 */ {0x18, false, nil, regTest{10: 0x84}, cacheTest{}},
			/*  70 */ {0x1c, false, nil, regTest{11: 0x55}, cacheTest{}},
			/*  71 */ {0x20, false, nil, regTest{}, cacheTest{0x100: 0x55}},
			/*  72 */ {0x24, false, nil, regTest{13: 0x104}, cacheTest{}},
			/*  73 */ {0x28, false, nil, regTest{6: 10}, cacheTest{}},
			/*  74 */ {0x2c, false, nil, regTest{8: 0}, cacheTest{}},
			/*  75 */ {0x30, false, nil, regTest{}, cacheTest{}},
			/*  76 */ {0x34, false, nil, regTest{6: 0}, cacheTest{}},
			/*  77 */ {0x38, false, nil, regTest{9: 0xdc}, cacheTest{}},
			/*  78 */ {0x3c, false, nil, regTest{7: 0x6}, cacheTest{}},
			/*  79 */ {0x40, false, nil, regTest{0: 0x6}, cacheTest{}},
			/*  80 */ {0x44, false, nil, regTest{6: 1}, cacheTest{}},
			/*  81 */ {0x48, false, nil, regTest{9: 0xe0}, cacheTest{}},
			/*  82 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  83 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/*  84 */ {0x3c, false, nil, regTest{7: 0x2c}, cacheTest{}},
			/*  85 */ {0x40, false, nil, regTest{0: 0x32}, cacheTest{}},
			/*  86 */ {0x44, false, nil, regTest{6: 2}, cacheTest{}},
			/*  87 */ {0x48, false, nil, regTest{9: 0xe4}, cacheTest{}},
			/*  88 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  89 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/*  90 */ {0x3c, false, nil, regTest{7: 0x45}, cacheTest{}},
			/*  91 */ {0x40, false, nil, regTest{0: 0x77}, cacheTest{}},
			/*  92 */ {0x44, false, nil, regTest{6: 3}, cacheTest{}},
			/*  93 */ {0x48, false, nil, regTest{9: 0xe8}, cacheTest{}},
			/*  94 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/*  95 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/*  96 */ {0x3c, false, nil, regTest{7: 0x1}, cacheTest{}},
			/*  97 */ {0x40, false, nil, regTest{0: 0x78}, cacheTest{}},
			/*  98 */ {0x44, false, nil, regTest{6: 4}, cacheTest{}},
			/*  99 */ {0x48, false, nil, regTest{9: 0xec}, cacheTest{}},
			/* 100 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 101 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 102 */ {0x3c, false, nil, regTest{7: 0x7}, cacheTest{}},
			/* 103 */ {0x40, false, nil, regTest{0: 0x7f}, cacheTest{}},
			/* 104 */ {0x44, false, nil, regTest{6: 5}, cacheTest{}},
			/* 105 */ {0x48, false, nil, regTest{9: 0xf0}, cacheTest{}},
			/* 106 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 107 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 108 */ {0x3c, false, nil, regTest{7: 0}, cacheTest{}},
			/* 109 */ {0x40, false, nil, regTest{0: 0x7f}, cacheTest{}},
			/* 110 */ {0x44, false, nil, regTest{6: 6}, cacheTest{}},
			/* 111 */ {0x48, false, nil, regTest{9: 0xf4}, cacheTest{}},
			/* 112 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 113 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 114 */ {0x3c, false, nil, regTest{7: 1}, cacheTest{}},
			/* 115 */ {0x40, false, nil, regTest{0: 0x80}, cacheTest{}},
			/* 116 */ {0x44, false, nil, regTest{6: 7}, cacheTest{}},
			/* 117 */ {0x48, false, nil, regTest{9: 0xf8}, cacheTest{}},
			/* 118 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 119 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 120 */ {0x3c, false, nil, regTest{7: 0x5}, cacheTest{}},
			/* 121 */ {0x40, false, nil, regTest{0: 0x85}, cacheTest{}},
			/* 122 */ {0x44, false, nil, regTest{6: 8}, cacheTest{}},
			/* 123 */ {0x48, false, nil, regTest{9: 0xfc}, cacheTest{}},
			/* 124 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 125 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 126 */ {0x3c, false, nil, regTest{7: 0xa}, cacheTest{}},
			/* 127 */ {0x40, false, nil, regTest{0: 0x8f}, cacheTest{}},
			/* 128 */ {0x44, false, nil, regTest{6: 9}, cacheTest{}},
			/* 129 */ {0x48, false, nil, regTest{9: 0x100}, cacheTest{}},
			/* 130 */ {0x4c, false, nil, regTest{8: 1}, cacheTest{}},
			/* 131 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 132 */ {0x3c, false, nil, regTest{7: 0x55}, cacheTest{}},
			/* 133 */ {0x40, false, nil, regTest{0: 0xe4}, cacheTest{}},
			/* 134 */ {0x44, false, nil, regTest{6: 10}, cacheTest{}},
			/* 135 */ {0x48, false, nil, regTest{9: 0x104}, cacheTest{}},
			/* 136 */ {0x4c, false, nil, regTest{8: 0}, cacheTest{}},
			/* 137 */ {0x50, false, nil, regTest{}, cacheTest{}},
			/* 138 */ {0x54, false, nil, regTest{}, cacheTest{0xac: 0xe4}},
			/* 139 */ {0x58, true, nil, regTest{}, cacheTest{}},
		},
	},
}

func TestCore(t *testing.T) {
	for _, tt := range coreTests {
		c := Mock(process.Mock(tt.sampleProgram))

		for i, tp := range tt.stateTestSequence {
			currentState := c.Process.State()
			callsign := fmt.Sprintf("[%d:%d]", c.Process.ProcessNumber, i)

			// set up next state
			c.Next = currentState.Next()

			// run the appropriate instruction
			c.Call()

			// test program counter
			pc := currentState.ProgramCounter
			if pc != tp.pc {
				t.Errorf(
					"%s PC: %04X (expected %04X)",
					callsign, uint32(pc), uint32(tp.pc),
				)
			}

			// test halt state
			if c.Next.Halt != tp.halt {
				t.Errorf(
					"%s Halt: %v (expected %v)",
					callsign, c.Next.Halt, tp.halt,
				)
			}

			// test error state
			if c.Next.Error != tp.err {
				t.Errorf(
					"%s Error: %v (expected %v)",
					callsign, c.Next.Error, tp.err,
				)
			}

			// test registers
			for rd, rv := range tp.registers {
				regActual := c.Next.Registers[rd]
				if regActual != rv {
					t.Errorf(
						"%s %v: %v (expected %v)",
						callsign, rd, regActual, rv,
					)
				}
			}

			// test caches
			for addr, wdval := range tp.caches {
				wdActual := c.Next.Caches.AddressFetchWord(addr)
				if wdActual != wdval {
					t.Errorf(
						"%s %v: %v (expected %v)",
						callsign, addr, wdActual, wdval,
					)
				}
			}

			// apply the next state
			c.Process.SetState(currentState.Apply(c.Next))
		}
	}
}

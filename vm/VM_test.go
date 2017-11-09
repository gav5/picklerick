package vm

import (
	"reflect"
	"testing"

	"./ivm"
)

func TestVMInstructionProxy(t *testing.T) {
	vm := VM{
		Clock: 42,
		Cores: [NumCores]Core{
			Core{
				PC: 42,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0xFEEDFACE, 0xDEADBEEF, 0x12345678,
				},
			},
			Core{
				PC: 1,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0x12345678, 0xFEEDFACE, 0xDEADBEEF,
				},
			},
			Core{
				PC: 2,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0xDEADBEEF, 0x12345678, 0xFEEDFACE,
				},
			},
			Core{
				PC: 3,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0x87654321, 0x11111111, 0x22222222,
				},
			},
		},
		RAM: RAM{
			0x00000000, 0x11111111, 0x22222222, 0x33333333,
			0x44444444, 0x55555555, 0x66666666, 0x77777777,
			0x88888888, 0x99999999, 0xaaaaaaaa, 0xbbbbbbbb,
			0xcccccccc, 0xdddddddd, 0xeeeeeeee, 0xffffffff,
		},
		Disk: Disk{},
	}
	ip1 := vm.instructionProxy(0)
	ip2 := ivm.MakeInstructionProxy(&vm.Cores[0], &vm.RAM)
	if !reflect.DeepEqual(ip1, ip2) {
		t.Errorf("ip1 != ip2; expected %v to equal %v", ip1, ip2)
	}
}

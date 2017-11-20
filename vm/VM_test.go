package vm

import (
	"reflect"
	"testing"

	"./ivm"
	"./core"
)

func TestVMIVMRelationship(t *testing.T) {
	vm := &VM{}
	_ = ivm.IVM(vm)
}

func TestVMInstructionProxy(t *testing.T) {
	vm := VM{
		Clock: 42,
		Cores: [ivm.NumCores]*core.Core{
			&core.Core{
				PC: 42,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0xFEEDFACE, 0xDEADBEEF, 0x12345678,
				},
			},
			&core.Core{
				PC: 1,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0x12345678, 0xFEEDFACE, 0xDEADBEEF,
				},
			},
			&core.Core{
				PC: 2,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0xDEADBEEF, 0x12345678, 0xFEEDFACE,
				},
			},
			&core.Core{
				PC: 3,
				Registers: [ivm.NumCoreRegisters]ivm.Word{
					0x87654321, 0x11111111, 0x22222222,
				},
			},
		},
		RAM: RAM{
			ivm.Frame{0x00000000, 0x11111111, 0x22222222, 0x33333333},
			ivm.Frame{0x44444444, 0x55555555, 0x66666666, 0x77777777},
			ivm.Frame{0x88888888, 0x99999999, 0xaaaaaaaa, 0xbbbbbbbb},
			ivm.Frame{0xcccccccc, 0xdddddddd, 0xeeeeeeee, 0xffffffff},
			ivm.Frame{0x00000000, 0x11111111, 0x22222222, 0x33333333},
			ivm.Frame{0x44444444, 0x55555555, 0x66666666, 0x77777777},
			ivm.Frame{0x88888888, 0x99999999, 0xaaaaaaaa, 0xbbbbbbbb},
			ivm.Frame{0xcccccccc, 0xdddddddd, 0xeeeeeeee, 0xffffffff},
		},
		Disk: Disk{},
	}
	ip1 := vm.InstructionProxy(vm.Cores[0])
	ip2 := ivm.MakeInstructionProxy(vm.Cores[0], &vm.RAM)
	if !reflect.DeepEqual(ip1, ip2) {
		t.Errorf("ip1 != ip2; expected %v to equal %v", ip1, ip2)
	}
}

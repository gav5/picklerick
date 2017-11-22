package vm

import (
	// "reflect"
	"testing"

	"./ivm"
	// "./core"
)

func TestVMIVMRelationship(t *testing.T) {
	vm := &VM{}
	_ = ivm.IVM(vm)
}

// func TestVMInstructionProxy(t *testing.T) {
// 	vm := New(config.Config{}){
// 		Clock: 42,
// 		Cores: [ivm.NumCores]*core.Core{},
// 		RAM: RAM{
// 			ivm.Frame{0x00000000, 0x11111111, 0x22222222, 0x33333333},
// 			ivm.Frame{0x44444444, 0x55555555, 0x66666666, 0x77777777},
// 			ivm.Frame{0x88888888, 0x99999999, 0xaaaaaaaa, 0xbbbbbbbb},
// 			ivm.Frame{0xcccccccc, 0xdddddddd, 0xeeeeeeee, 0xffffffff},
// 			ivm.Frame{0x00000000, 0x11111111, 0x22222222, 0x33333333},
// 			ivm.Frame{0x44444444, 0x55555555, 0x66666666, 0x77777777},
// 			ivm.Frame{0x88888888, 0x99999999, 0xaaaaaaaa, 0xbbbbbbbb},
// 			ivm.Frame{0xcccccccc, 0xdddddddd, 0xeeeeeeee, 0xffffffff},
// 		},
// 		Disk: Disk{},
// 	}
// 	for num := uint8(0); num < ivm.NumCores; num++ {
// 		vm.Cores[num] = core.New(num)
// 	}
// 	ip1 := vm.InstructionProxy(vm.Cores[0])
// 	ip2 := ivm.MakeInstructionProxy(
// 		vm.Cores[0], &vm.RAM, vm.Cores[0].PagingProxy(),
// 	)
// 	if !reflect.DeepEqual(ip1, ip2) {
// 		t.Errorf("ip1 != ip2; expected %v to equal %v", ip1, ip2)
// 	}
// }

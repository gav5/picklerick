package vm

import (
	"reflect"
	"testing"

	"./ivm"
)

func TestRAMWord(t *testing.T) {
	ram := RAM{}
	v1 := ivm.Word(0xDEADBEEF)
	ram.AddressWriteWord(0x42, v1)
	v2 := ram.AddressFetchWord(0x42)

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("v1 != v2; expected %v to equal %v", v1, v2)
	}
}

func TestRAMUint32(t *testing.T) {
	ram := RAM{}
	v1 := uint32(0xFEEDFACE)
	ram.AddressWriteUint32(0x12, v1)
	v2 := ram.AddressFetchUint32(0x12)

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("v1 != v2; expected %v to equal %v", v1, v2)
	}
}

func TestRAMInt32(t *testing.T) {
	ram := RAM{}
	v1 := int32(-42)
	ram.AddressWriteInt32(0x55, v1)
	v2 := ram.AddressFetchInt32(0x55)

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("v1 != v2; expected %v to equal %v", v1, v2)
	}
}

func TestRAMBool(t *testing.T) {
	ram := RAM{}
	v1 := true
	ram.AddressWriteBool(0x77, v1)
	v2 := ram.AddressFetchBool(0x77)

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("v1 != v2; expected %v to equal %v", v1, v2)
	}
}

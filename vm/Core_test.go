package vm

import (
	"reflect"
	"testing"

	"./ivm"
)

func TestCoreProgramCounter(t *testing.T) {
	c := &Core{}
	c1 := ivm.Address(42)
	c.SetProgramCounter(c1)
	c2 := c.ProgramCounter()

	if !reflect.DeepEqual(c1, c2) {
		t.Errorf("c1 != c2; expected %v to equal %v", c1, c2)
	}
}

func TestCoreRegisterWord(t *testing.T) {
	c := &Core{}
	r1 := ivm.Word(0x12345678)
	c.SetRegisterWord(0x0, r1)
	r2 := c.RegisterWord(0x0)

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("r1 != r2; expected %v to equal %v", r1, r2)
	}
}

func TestCoreRegisterUint32(t *testing.T) {
	c := &Core{}
	r1 := uint32(0x12345678)
	c.SetRegisterUint32(0x0, r1)
	r2 := c.RegisterUint32(0x0)

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("r1 != r2; expected %v to equal %v", r1, r2)
	}
}

func TestCoreRegisterInt32(t *testing.T) {
	c := &Core{}
	r1 := int32(-42)
	c.SetRegisterInt32(0x0, r1)
	r2 := c.RegisterInt32(0x0)

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("r1 != r2; expected %v to equal %v", r1, r2)
	}
}

func TestCoreRegisterBool(t *testing.T) {
	c := &Core{}
	r1 := true
	c.SetRegisterBool(0x0, r1)
	r2 := c.RegisterBool(0x0)

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("r1 != r2; expected %v to equal %v", r1, r2)
	}
}

package kernel

import (
  "../vm/ivm"
  "../prog"
  "../config"
  "log"
)

// Kernel houses all the storage and functionality of the OS kernel.
type Kernel struct {
  config config.Config
  virtualMachine ivm.IVM
  processTable processTableType
  frameTable frameTableType
}

// MakeKernel makes a kernel with the given virtual machine.
func MakeKernel(virtualMachine ivm.IVM, c config.Config) (Kernel, error) {
  k := Kernel{
    config: c,
    virtualMachine: virtualMachine,
    processTable: processTableType{},
    frameTable: frameTableType{},
  }
  // load programs into the system
  var programArray []prog.Program
  var err error
	if programArray, err = prog.ParseFile(c.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return k, err
	}
  k.LoadPrograms(programArray)
  return k, nil
}

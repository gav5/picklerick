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
  ramFrameTable frameTableType
  diskFrameTable frameTableType
}

// MakeKernel makes a kernel with the given virtual machine.
func MakeKernel(virtualMachine ivm.IVM, c config.Config) (Kernel, error) {
  k := Kernel{
    config: c,
    virtualMachine: virtualMachine,
    processTable: processTableType{},
    ramFrameTable: frameTableType{},
    diskFrameTable: frameTableType{},
  }
  // load programs into the system
  var programArray []prog.Program
  var err error
	if programArray, err = prog.ParseFile(c.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return k, err
	}
  log.Printf("Got %d programs!\n", len(programArray))
  if err = k.LoadPrograms(programArray); err != nil {
    return k, err
  }
  return k, nil
}

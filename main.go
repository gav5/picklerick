package main

import (
	"log"

	"./config"
	"./vm"
	"os"
	"fmt"
	// "./disp"
)

func main() {
	var err error
	var sharedConfig config.Config
	var virtualMachine *vm.VM

	// load app configuration
	sharedConfig, err = config.Shared()
	if err != nil {
		log.Fatalf("error extracting shared configuration: %v", err)
		return
	}

	// introduce program, display configuration
	log.Println("picklerick OS")
	log.Printf("program file: %s\n", sharedConfig.Progfile)

	// build the virtual machine with the given config
	virtualMachine, err = vm.New(sharedConfig)
	if err != nil {
		log.Fatalf("error building virtual machine: %v", err)
		return
	}
	virtualMachine.Tick()

	fmt.Println()
	_ = virtualMachine.FprintProcessTable(os.Stdout)

	// fmt.Println()
	// virtualMachine.Run()

	fmt.Print("\nRAM Dump:\n")
	_ = virtualMachine.RAM.Print()
	fmt.Print("\n")

	fmt.Print("\nDisk Dump:\n")
	_ = virtualMachine.Disk.Print()
	fmt.Print("\n")
}

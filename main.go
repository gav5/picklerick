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
	fmt.Println("\npicklerick OS")
	fmt.Printf("progfile:\t%s\n", sharedConfig.Progfile)
	fmt.Printf("outdir:\t\t%s\n", sharedConfig.Outdir)
	fmt.Printf("sched:\t\t%s\n", sharedConfig.Sched)
	fmt.Printf("qsize:\t\t%d\n", sharedConfig.QSize)
	fmt.Println()

	// build the virtual machine with the given config
	virtualMachine, err = vm.New(sharedConfig)
	if err != nil {
		log.Fatalf("error building virtual machine: %v", err)
		return
	}
	virtualMachine.Tick()

	fmt.Println()
	virtualMachine.FprintProcessTable(os.Stdout)

	fmt.Println("\nExecution Logs")
	virtualMachine.Run()

	// fmt.Println()
	// _ = virtualMachine.FprintProcessTable(os.Stdout)

	// fmt.Print("\nRAM Dump:\n")
	// virtualMachine.RAM.Print()
	// fmt.Print("\n")

	// fmt.Print("\nDisk Dump:\n")
	// _ = virtualMachine.Disk.Print()
	// fmt.Print("\n")
}

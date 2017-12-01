package main

import (
	"fmt"
	"log"

	"./config"
	"./report"
	"./util/logger"
	"./vm"
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

	// set up logger based on given configurations
	if err = logger.Init(sharedConfig); err != nil {
		log.Fatalf("error setting up logger: %v", err)
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
		log.Fatalf("ERROR building VM: %v", err)
		return
	}

	// run the virtual machine
	err = virtualMachine.Run()
	if err != nil {
		// if this fails, we want to keep going
		// (this is because we still want reports to be generated)
		log.Printf("ERROR running VM: %v", err)
	}

	// build the reports and save to disk
	err = report.Generate(sharedConfig, virtualMachine)
	if err != nil {
		log.Fatalf("ERROR generating reports: %v", err)
	}

	// let the user know this finished successfully
	fmt.Println("\ndone!")
}

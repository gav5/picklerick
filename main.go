package main

import (
	"log"

	"./config"
	"./vm"
	"./disp"
)

func main() {
	var err error
	var sharedConfig config.Config
	var virtualMachine vm.VM

	// load app configuration
	sharedConfig, err = config.Shared()
	if err != nil {
		log.Fatalf("error extracting shared configuration: %v", err)
		return
	}

	// introduce program, display configuration
	log.Println("picklerick OS")
	log.Printf("program file: %s\n", sharedConfig.Progfile)
	fmt.Println()

	// build the virtual machine with the given config
	virtualMachine, err = vm.MakeVM(sharedConfig)
	if err != nil {
		log.Fatalf("error building virtual machine: %v", err)
		return
	}

	fmt.Print("RAM Dump:\n")
	_ = virtualMachine.RAM.Print()
	fmt.Print("\n")

	fmt.Print("Disk Dump:\n")
	_ = virtualMachine.Disk.Print()
	fmt.Print("\n")

	// disp.RunTask(virtualMachine)

	// c := cpu.State{}

	// if len(sharedConfig.Outdir) > 0 {
	// 	disp.CleanOutDir(sharedConfig.Outdir)
	// }

	// for _, p := range programArray {
	// 	pcb := proc.MakePCB(p)
	// 	c.ContextSwitch(pcb)
	// 	for !c.ShouldHalt {
	// 		c.Next()
	// 	}
	// 	if len(sharedConfig.Outdir) > 0 {
	// 		disp.ProgramOutputFile(sharedConfig.Outdir, cpuType.CPU{ID: 1, State: cpuType.State(c)})
	// 	}
	// }

	// load the ASM output file (if applicable)
	// if len(sharedConfig.Outdir) > 0 {
	// 	disp.MakeAll(sharedConfig.Outdir, programArray, []cpuType.CPU{
	// 		cpuType.CPU{ID: 1, State: cpuType.State(c)},
	// 	})
	// }
}

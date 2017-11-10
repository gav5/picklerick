package main

import (
	"fmt"
	"log"

	"./config"
	"./kernel"
	"./prog"
	"./vm"
)

func main() {
	var err error
	var sharedConfig config.Config
	var programArray []prog.Program

	// load app configuration
	if sharedConfig, err = config.Shared(); err != nil {
		log.Fatalf("error extracting shared configuration: %v", err)
		return
	}

	// parse file and display assembly code for each job
	if programArray, err = prog.ParseFile(sharedConfig.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return
	}

	// introduce program, display configuration
	log.Println("picklerick OS")
	log.Printf("program file: %s\n", sharedConfig.Progfile)
	fmt.Println()

	virtualMachine := vm.MakeVM()
	kernel.LoadPrograms(&virtualMachine, programArray)

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

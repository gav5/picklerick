package main

import (
	"log"

	"./config"
	"./cpu"
	"./disp"
	"./prog"
	"./ram"
)

func main() {
	ram.SetData(0x01, 0xFFFFFFFF)
	ram.SetData(0x02, 0xAAAAAAAA)

	var err error
	var sharedConfig config.Config
	var programArray []prog.Program

	// load app configuration
	if sharedConfig, err = config.Shared(); err != nil {
		log.Fatalf("error extracting shared configuration: %v", err)
		return
	}

	// introduce program, display configuration
	log.Println("picklerick OS")
	log.Printf("program file: %s\n", sharedConfig.Progfile)

	// parse file and display assembly code for each job
	if programArray, err = prog.ParseFile(sharedConfig.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return
	}

	cpus := []cpu.CPU{
		cpu.CPU{ID: 1, State: cpu.State{}},
	}

	// load the ASM output file (if applicable)
	if len(sharedConfig.Outdir) > 0 {
		disp.MakeAll(sharedConfig.Outdir, programArray, cpus)
	}
}

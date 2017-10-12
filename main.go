package main

import (
	"log"

	"./config"
	"./prog"
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

	// introduce program, display configuration
	log.Println("picklerick OS")
	log.Printf("program file: %s\n", sharedConfig.Progfile)

	// parse file and display assembly code for each job
	if programArray, err = prog.ParseFile(sharedConfig.Progfile); err != nil {
		log.Fatalf("error parsing program file: %v\n", err)
		return
	}

	// load the ASM output file (if applicable)
	if len(sharedConfig.ASMFile) > 0 {
		log.Printf("ASM output file: %s", sharedConfig.ASMFile)
		f, err := prog.MakeASMFile(sharedConfig.ASMFile)
		if err != nil {
			log.Fatalf("error opening ASM file: %v\n", err)
			return
		}
		if err = f.WritePrograms(programArray, sharedConfig.Progfile); err != nil {
			log.Fatalf("error writing to ASM file: %v", err)
		}
		if err = f.Close(); err != nil {
			log.Fatalf("error closing ASM file: %v\n", err)
		}
	}
}

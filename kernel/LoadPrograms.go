package kernel

import (
	// "log"
	"../prog"
	"./page"
	"./process"
)

// LoadPrograms loads the given programs into the virtual machine.
func (k *Kernel) LoadPrograms(programs []prog.Program) error {
	for _, p := range programs {
		if err := k.loadProgram(p); err != nil {
			return err
		}
	}
	return nil
}

func (k *Kernel) loadProgram(program prog.Program) error {
	// log.Printf("Loading program %d...\n", program.Job.ID)
	// add the given pages into the VM (starting with RAM, then disk)
	pageTable := page.Table{}
	if err := k.PushProgram(program, &pageTable); err != nil {
		return err
	}
	// add to the process manager
	k.pm.Add(process.Make(program, pageTable))
	return nil
}

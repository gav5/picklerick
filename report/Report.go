package report

import (
	"fmt"

	"../config"
	"../vm"
)

var sharedConfig = config.Default

// Generate makes the necessary reports for the given vm and saves.
func Generate(cfg config.Config, virtualMachine *vm.VM) error {
	sharedConfig = cfg
	fmt.Println("Report Generation:")

	type reportingInstance = struct {
		r reportBase
		b reportBuilder
	}
	reportMakersArray := []reportMaker{
		makeRAMDumpReport,
		makeDiskDumpReport,
		makeProcessTableReport,
		makeProcessResultReports,
	}

	for _, rm := range reportMakersArray {
		rary, b, err := rm(virtualMachine)
		if err != nil {
			return err
		}
		// go through each report from the list
		for _, r := range rary {
			fname, err := saveReportFile(r, b)
			if err != nil {
				fmt.Printf("✗ ERROR generating %s: %v\n", fname, err)
			} else {
				fmt.Printf("✓ ADDED %s\n", fname)
			}
		}
	}
	return nil
}

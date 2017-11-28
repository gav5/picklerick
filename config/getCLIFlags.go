package config

import "flag"

const (
	progfileDesc  = "program file for system to execute"
	outdirDesc    = "directory used to store output data"
	schedDesc		  = "scheduler method to use"
	qsizeDesc		  = "the size used to hold the long-term queue"
	maxCyclesDesc = "the maximum VM cycles before breaking"
)

func getCLIFlags(config *Config) error {
	flag.StringVar(&globalShared.Progfile, "progfile", globalShared.Progfile, progfileDesc)
	flag.StringVar(&globalShared.Outdir, "outdir", globalShared.Outdir, outdirDesc)
	flag.StringVar(&globalShared.Sched, "sched", globalShared.Sched, schedDesc)
	flag.UintVar(&globalShared.QSize, "qsize", globalShared.QSize, qsizeDesc)
	flag.UintVar(&globalShared.MaxCycles, "maxcycles", globalShared.MaxCycles, maxCyclesDesc)
	flag.Parse()
	return nil
}

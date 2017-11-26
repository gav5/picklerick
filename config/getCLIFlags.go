package config

import "flag"

const (
	progfileDesc = "program file for system to execute"
	outdirDesc   = "directory used to store output data"
	schedDesc		 = "scheduler method to use"
)

func getCLIFlags(config *Config) error {
	flag.StringVar(&globalShared.Progfile, "progfile", globalShared.Progfile, progfileDesc)
	flag.StringVar(&globalShared.Outdir, "outdir", globalShared.Outdir, outdirDesc)
	flag.StringVar(&globalShared.Sched, "sched", globalShared.Sched, schedDesc)
	flag.Parse()
	return nil
}

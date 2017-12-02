package config

import "flag"

const (
	progfileDesc  = "program file for system to execute"
	outdirDesc    = "directory used to store output data"
	schedDesc     = "scheduler method to use"
	qsizeDesc     = "the size used to hold the long-term queue"
	maxCyclesDesc = "the maximum VM cycles before breaking"
	quietDesc     = "supress module logging to Stdout"
	dumpAtDesc    = "what processes/PCs to dump at (ex: 2@4C,3@04)"
)

func getCLIFlags(config *Config) error {
	flag.StringVar(
		&globalShared.Progfile, "progfile",
		globalShared.Progfile, progfileDesc,
	)
	flag.StringVar(
		&globalShared.Outdir, "outdir",
		globalShared.Outdir, outdirDesc,
	)
	flag.StringVar(
		&globalShared.Sched, "sched",
		globalShared.Sched, schedDesc,
	)
	flag.UintVar(
		&globalShared.QSize, "qsize",
		globalShared.QSize, qsizeDesc,
	)
	flag.UintVar(
		&globalShared.MaxCycles, "maxcycles",
		globalShared.MaxCycles, maxCyclesDesc,
	)
	flag.BoolVar(
		&globalShared.Quiet, "quiet",
		globalShared.Quiet, quietDesc,
	)
	flag.StringVar(
		&globalShared.DumpAt, "dumpat",
		globalShared.DumpAt, dumpAtDesc,
	)
	flag.Parse()
	return nil
}

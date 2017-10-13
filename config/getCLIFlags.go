package config

import "flag"

const (
	progfileDesc = "program file for system to execute"
	outdirDesc   = "directory used to store output data"
)

func getCLIFlags(config *Config) error {
	flag.StringVar(&globalShared.Progfile, "progfile", globalShared.Progfile, progfileDesc)
	flag.StringVar(&globalShared.Outdir, "outdir", globalShared.Outdir, outdirDesc)
	flag.Parse()
	return nil
}

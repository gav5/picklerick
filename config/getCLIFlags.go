package config

import "flag"

const (
	progfileDesc = "program file for system to execute"
	asmfileDesc  = "assembly file for system to output to (blank doesn't write to a file)"
)

func getCLIFlags(config *Config) error {
	flag.StringVar(&globalShared.Progfile, "progfile", globalShared.Progfile, progfileDesc)
	flag.StringVar(&globalShared.ASMFile, "asmfile", globalShared.ASMFile, asmfileDesc)
	flag.Parse()
	return nil
}

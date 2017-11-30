package config

// Config describes the configration of the application
type Config struct {
	Progfile  string `json:"progfile"`
	Outdir    string `json:"outdir"`
	Sched     string `json:"sched"`
	QSize     uint   `json:"qsize"`
	MaxCycles uint   `json:"maxcycles"`
	Quiet     bool   `json:"quiet"`
	DumpAt    string `json:"dump_at"`
}

// Default is the default value to use for d/c conditions
var Default = Config{
	Progfile:  "Program-File.txt",
	Outdir:    "out",
	Sched:     "fcfs",
	QSize:     200,
	MaxCycles: 100000,
	Quiet:     true,
	DumpAt:    "",
}

// private global variable for the config value
var globalShared *Config

// Shared gets the shared application config value
func Shared() (Config, error) {
	if globalShared == nil {
		// globalShared is not initialized yet
		// (so we should initialize it with the appropriate values)
		globalShared = &Config{}

		// first, let's load the appropriate defaults from file
		err := getRCFile(globalShared)
		if err != nil {
			return Config{}, err
		}

		// next, let's parse the flags from the command-line
		err = getCLIFlags(globalShared)
		if err != nil {
			return Config{}, err
		}
	}
	return *globalShared, nil
}

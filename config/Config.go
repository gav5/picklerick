package config

// Config describes the configration of the application
type Config struct {
	Progfile string `json:"progfile"`
	Outdir   string `json:"outdir"`
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

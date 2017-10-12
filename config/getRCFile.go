package config

import (
	"io/ioutil"

	"encoding/json"
)

func getRCFile(config *Config) error {
	file, err := ioutil.ReadFile(".picklerickrc.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		return err
	}
	return nil
}

package util

import (
	"errors"
	"io/ioutil"

	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/config"
	yaml "gopkg.in/yaml.v3"
)

var loadedConfigs *config.Server

// TODO: Config loader. Must read config from yaml file or os inv!
func loadAll() error {
	configs := config.Server{}

	// loading config.yml
	yamlFile, err := ioutil.ReadFile(config.Config_Path)
	if err != nil {
		return errors.New("failed to read config file! reason: " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		return errors.New("failed to unmarshal yml file! reason: " + err.Error())
	}

	loadedConfigs = &configs
	return nil
}

// Return config file if exists!
func GetConfigs() (config.Server, error) {
	var err error = nil
	if loadedConfigs == nil {
		err = loadAll()
	}
	if err != nil {
		return config.Server{}, err
	}
	conf := *loadedConfigs
	return conf, nil
}

package util

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/config"
	yaml "gopkg.in/yaml.v3"
)

var loadedConfigs *config.Server

// Config loader. Must read config from yaml file or os inv!
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

	// Read os env vars if exists
	redis_addr := os.Getenv("REDIS_ADDR")
	if redis_addr != "" {
		configs.REDIS_ADDR = redis_addr
	}
	api_key := os.Getenv("API_KEY")
	if api_key != "" {
		configs.API_KEY = api_key
	}
	rebrandlyUrl := os.Getenv("REBRANDLY_URL")
	if rebrandlyUrl != "" {
		configs.API_KEY = rebrandlyUrl
	}
	redis_time := os.Getenv("REDIS_TIME")
	if rebrandlyUrl != "" {
		rt, err := strconv.Atoi(redis_time)
		if err != nil {
			configs.REDIS_TIME = 300 // 5 min is default
		} else {
			configs.REDIS_TIME = rt
		}
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

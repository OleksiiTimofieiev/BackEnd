package configs

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// BackEnd []string
	Work string
}

const (
	configFile = "/home/user/Desktop/BackEnd/Go/src/GoogleDriveUpload/configs/configs.json"
)

func ReadConfigs(config *Config) {
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
}

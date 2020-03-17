package client

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config is Config
type Config struct {
	Hub struct {
		Address string
		Port    string
	}

	Self struct {
		Address string
		Port    string
	}
}

// IConfig is Config Single
var IConfig Config
var once sync.Once

// GetConfig .
func GetConfig() *Config {
	once.Do(analysisConfigYaml)
	return &IConfig
}

func analysisConfigYaml() {

	// TODO: Change
	data, err := ioutil.ReadFile("../tclient.yaml")
	str := string(data)
	println(str)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(data, &IConfig); err != nil {
		panic(err)
	}
}

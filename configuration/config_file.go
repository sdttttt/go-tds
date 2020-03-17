package configuration

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

// config is config
type config struct {
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
var IConfig config
var once sync.Once

// GetConfig .
func GetConfig() *config {
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

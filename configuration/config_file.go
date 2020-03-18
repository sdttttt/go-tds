package configuration

import (
	"io/ioutil"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// IConfig is Config Single
var IConfig config
var once sync.Once

var configPath string = ""

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

// GetConfig .
func GetConfig() *config {
	once.Do(analysisConfigYaml)
	return &IConfig
}

// ChangeConfigFilePath is Change Configuration File Path
func ChangeConfigFilePath(path string) {
	configPath = path
}

func analysisConfigYaml() {

	var data []byte
	var err error

	if configPath != "" {
		data, err = ioutil.ReadFile(configPath)
	} else {
		data, err = readDefaultPathConfigfile()
	}

	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(data, &IConfig); err != nil {
		panic(err)
	}
}

func readDefaultPathConfigfile() ([]byte, error) {

	var data []byte
	var err error

	paths := strings.Split(DefaultConfigPath, ",")

	for _, path := range paths {
		data, err = ioutil.ReadFile(path)
		if err == nil {
			break
		}
	}

	return data, err
}

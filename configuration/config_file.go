package configuration

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// IConfig is Config Single
var IConfig Config
var once sync.Once

var configPath string = ""

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

// GetConfig .
func GetConfig() *Config {
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
		log.Fatalln(err)
		return
	}
	if err = yaml.Unmarshal(data, &IConfig); err != nil {
		log.Fatalln(err)
		return
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

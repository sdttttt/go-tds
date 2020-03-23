package configuration

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

// IConfig is Config Single
var IConfig Config

// Single Instance.
var once sync.Once

// Custom definition Config file path.
var configPath string = ""

// Config is Config
type Config struct {
	Hub struct {
		Address           string
		Port              string
		CheckSurvivalTime time.Duration `yaml:"checkSurvivalTime"`
	}

	Self struct {
		Address      string
		Port         string
		SurvivalTime time.Duration `yaml:"survivalTime"`
	}
}

// GetConfig is to Read config.
func GetConfig() *Config {
	once.Do(analysisConfigYaml)
	return &IConfig
}

// Refresh description: Configuration is a singleton.
// If you need to change the configuration file path in the middle
// please use `Review` after using `ChangeConfigFilePath`.
func Refresh() {
	analysisConfigYaml()
}

// ChangeConfigFilePath is Change Configuration File Path
func ChangeConfigFilePath(path string) {
	configPath = path
}

// analysisConfigYaml is yaml file to Golang Struct.
func analysisConfigYaml() {

	var data []byte
	var err error

	if configPath != "" {
		data, err = ioutil.ReadFile(configPath)
	} else {
		data, err = readDefaultPathConfigfile()
	}

	if err != nil {
		log.Println(err.Error())
		return
	}
	if err = yaml.Unmarshal(data, &IConfig); err != nil {
		log.Println(err.Error())
		return
	}
}

// readDefaultPathConfigfile is read default file path.
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

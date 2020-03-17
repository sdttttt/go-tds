package configuration

import (
	"testing"
)

func TestConfigYaml(t *testing.T) {
	config := GetConfig()
	if config.Hub.Address == "" {
		t.Error("oh! not")
	}
	if config.Hub.Port == "" {
		t.Error("oh! not")
	}
}

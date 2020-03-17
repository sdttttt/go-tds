package tests

import (
	"testing"

	"github.com/sdttttt/go-tds/client"
)

func TestConfigYaml(t *testing.T) {
	config := client.GetConfig()
	if config.Hub.Address == "" {
		t.Error("oh! not")
	}
	if config.Hub.Port == "" {
		t.Error("oh! not")
	}
}

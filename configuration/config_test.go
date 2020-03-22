package configuration

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigYaml(t *testing.T) {

	Convey("Test YAML Configuration is Effective", t, func() {
		config := GetConfig()
		So(config, ShouldNotBeNil)

		Convey("Test YAML Configuration Attribute is Effective", func() {
			So(config.Hub.Address, ShouldNotBeBlank)
			So(config.Hub.Port, ShouldNotBeBlank)
			So(config.Hub.CheckSurvivalTime, ShouldNotBeZeroValue)
			So(config.Self.Address, ShouldNotBeBlank)
			So(config.Self.Port, ShouldNotBeBlank)
			So(config.Self.SurvivalTime, ShouldNotBeZeroValue)
		})
	})

	Convey("Test Read Default Path Config is Effective", t, func() {
		data, err := readDefaultPathConfigfile()

		So(data, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestChangeConfigYaml(t *testing.T) {

	ChangeConfigFilePath("../tclient.yaml")
	analysisConfigYaml()

	Convey("Test YAML Configuration is Effective", t, func() {
		config := GetConfig()
		So(config, ShouldNotBeNil)

		Convey("Test YAML Configuration Attribute is Effective", func() {
			So(config.Hub.Address, ShouldNotBeBlank)
			So(config.Hub.Port, ShouldNotBeBlank)
			So(config.Hub.CheckSurvivalTime, ShouldNotBeZeroValue)
			So(config.Self.Address, ShouldNotBeBlank)
			So(config.Self.Port, ShouldNotBeBlank)
			So(config.Self.SurvivalTime, ShouldNotBeZeroValue)
		})
	})

	Convey("Test Read Default Path Config is Effective", t, func() {
		data, err := readDefaultPathConfigfile()

		So(data, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestChangeConfigErrorYaml(t *testing.T) {
	ChangeConfigFilePath("../go.mod")
	analysisConfigYaml()
}

func TestErrorChangeConfigYaml(t *testing.T) {
	ChangeConfigFilePath("..")
	analysisConfigYaml()
}

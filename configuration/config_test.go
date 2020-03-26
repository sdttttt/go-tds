package configuration

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigYaml(t *testing.T) {

	Convey("Test YAML Configuration is Effective", t, func() {
		config := GetConfig()
		So(config, ShouldNotBeNil)

		hub := config.Hub
		self := config.Self

		Convey("Test YAML Configuration Attribute is Effective", func() {
			So(hub.Address, ShouldNotBeBlank)
			So(hub.Port, ShouldNotBeBlank)
			So(hub.CheckSurvivalTime, ShouldNotBeZeroValue)
			So(hub.LoadBalance, ShouldNotBeBlank)
			So(self.Address, ShouldNotBeBlank)
			So(self.Port, ShouldNotBeBlank)
			So(self.SurvivalTime, ShouldNotBeZeroValue)
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
	Refresh()

	Convey("Test YAML Configuration is Effective", t, func() {
		config := GetConfig()
		So(config, ShouldNotBeNil)

		hub := config.Hub
		self := config.Self

		Convey("Test YAML Configuration Attribute is Effective", func() {
			So(hub.Address, ShouldNotBeBlank)
			So(hub.Port, ShouldNotBeBlank)
			So(hub.CheckSurvivalTime, ShouldNotBeZeroValue)
			So(hub.LoadBalance, ShouldNotBeBlank)
			So(self.Address, ShouldNotBeBlank)
			So(self.Port, ShouldNotBeBlank)
			So(self.SurvivalTime, ShouldNotBeZeroValue)
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
	Refresh()
}

func TestErrorChangeConfigYaml(t *testing.T) {
	ChangeConfigFilePath("..")
	Refresh()
}

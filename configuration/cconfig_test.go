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
			So(config.Hub.Address, ShouldNotBeNil)
			So(config.Hub.Port, ShouldNotBeNil)
		})
	})

	Convey("Test Read Default Path Config is Effective", t, func() {
		data, err := readDefaultPathConfigfile()

		So(data, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

}

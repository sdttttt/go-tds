package main

import (
	"testing"
	"time"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/trpc"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHub(t *testing.T) {

	configuration.ChangeConfigFilePath("./tclient-test.yaml")

	// Run Hub
	go main()
	serviceName := "API.HelloWorld"

	// Wait Main Running
	time.Sleep(time.Duration(2) * time.Second)

	Convey("Test Error Get ServiceAddr", t, func() {
		// Service not Register
		info, err := trpc.GetServiceAddr(serviceName)
		So(err, ShouldBeNil)
		So(info.Ip, ShouldBeBlank)
		So(info.Port, ShouldBeBlank)
		So(info.ServiceName, ShouldNotBeBlank)
	})

	Convey("Test Error Call Service", t, func() {
		// Service not Register
		err := trpc.Call(serviceName, nil, nil)
		So(err, ShouldNotBeNil)
	})

	Convey("Test Service Register to Hub", t, func() {
		err := trpc.Register(serviceName)
		So(err, ShouldBeNil)
	})

	Convey("Test Get Service from Hub", t, func() {
		for i := 0; i < 2; i++ {
			info, err := trpc.GetServiceAddr(serviceName)
			So(err, ShouldBeNil)
			So(serviceName, ShouldEqual, info.ServiceName)
			So(info.GetIp(), ShouldNotBeBlank)
			So(info.GetPort(), ShouldNotBeBlank)
		}
	})

	time.Sleep(time.Duration(5) * time.Second)

	Convey("Test Error Get ServiceAddr", t, func() {
		// Service not Register
		info, err := trpc.GetServiceAddr(serviceName)
		So(err, ShouldBeNil)
		So(info.Ip, ShouldNotBeBlank)
		So(info.Port, ShouldNotBeBlank)
		So(info.ServiceName, ShouldNotBeBlank)
	})

	return
}

func TestHubTimeoutService(t *testing.T) {

	configuration.ChangeConfigFilePath("./tclient-test2.yaml")
	configuration.Refresh()

	// Run Hub
	go main()
	serviceName := "API.HelloWorld"

	// Wait Main Running
	time.Sleep(time.Duration(2) * time.Second)

	Convey("Test Error Get ServiceAddr", t, func() {
		// Service not Register
		info, err := trpc.GetServiceAddr(serviceName)
		So(err, ShouldBeNil)
		So(info.Ip, ShouldBeBlank)
		So(info.Port, ShouldBeBlank)
		So(info.ServiceName, ShouldNotBeBlank)
	})

	Convey("Test Error Call Service", t, func() {
		// Service not Register
		err := trpc.Call(serviceName, nil, nil)
		So(err, ShouldNotBeNil)
	})

	Convey("Test Service Register to Hub", t, func() {
		err := trpc.Register(serviceName)
		So(err, ShouldBeNil)
	})

	time.Sleep(time.Duration(2) * time.Second)

	Convey("Test Error Get ServiceAddr", t, func() {
		// Service not Register
		info, err := trpc.GetServiceAddr(serviceName)
		So(err, ShouldBeNil)
		So(info.Ip, ShouldBeBlank)
		So(info.Port, ShouldBeBlank)
		So(info.ServiceName, ShouldNotBeBlank)
	})

	return
}

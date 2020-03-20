package main

import (
	"testing"

	"github.com/sdttttt/go-tds/trpc"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHub(t *testing.T) {
	// Run Hub
	go main()
	serviceName := "API.HelloWorld"

	Convey("Test Error Get ServiceAddr", t, func() {
		info, err := trpc.GetServiceAddr(serviceName)
		So(err, ShouldBeNil)
		So(info.Ip, ShouldBeBlank)
		So(info.Port, ShouldBeBlank)
		So(info.ServiceName, ShouldNotBeBlank)
	})

	Convey("Test Error Call Service", t, func() {
		err := trpc.Call(serviceName, nil, nil)
		So(err, ShouldNotBeNil)
	})

	Convey("Test Service Register to Hub", t, func() {
		err := trpc.Register(serviceName)
		So(err, ShouldBeNil)
	})

	Convey("Test Get Service from Hub", t, func() {

		for i := 0; i < 10; i++ {
			info, err := trpc.GetServiceAddr(serviceName)
			So(err, ShouldBeNil)
			So(serviceName, ShouldEqual, info.ServiceName)
			So(info.GetIp(), ShouldNotBeBlank)
			So(info.GetPort(), ShouldNotBeBlank)
		}
	})

	return
}

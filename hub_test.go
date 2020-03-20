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
		ShouldBeNil(err)
		ShouldBeBlank(info.Ip)
		ShouldBeBlank(info.Port)
		ShouldNotBeBlank(info.GetServiceName())
	})

	Convey("Test Error Call Service", t, func() {
		err := trpc.Call(serviceName, nil, nil)
		ShouldNotBeNil(err)
	})

	Convey("Test Service Register to Hub", t, func() {
		err := trpc.Register(serviceName)
		ShouldBeNil(err)
	})

	Convey("Test Get Service from Hub", t, func() {

		for i := 0; i < 10; i++ {
			info, err := trpc.GetServiceAddr(serviceName)
			ShouldBeNil(err)
			ShouldEqual(serviceName, info.ServiceName)
			ShouldNotBeBlank(info.GetIp())
			ShouldNotBeBlank(info.GetPort())
		}

	})

	return
}

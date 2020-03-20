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

	Convey("Test Service Register to Hub", t, func() {
		err := trpc.Register(serviceName)
		ShouldBeNil(err)
	})

	Convey("Test Get Service from Hub", t, func() {
		info, err := trpc.GetServiceAddr(serviceName)
		ShouldBeNil(err)
		ShouldEqual(serviceName, info.ServiceName)
		ShouldNotBeBlank(info.GetIp())
		ShouldNotBeBlank(info.GetPort())
	})

	return
}

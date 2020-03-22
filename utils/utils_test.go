package utils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimer(t *testing.T) {

	Convey("Test Run Job", t, func() {
		timer := NewTimer()
		timer.AddJob(1, func() {
			print(" ")
		})
		timer.Start()
	})

	Convey("Test Run Jobs", t, func() {
		timer := NewTimer()
		timer.AddJobs(1, func() {
			print(" ")
		}, func() {
			print(" ")
		})
		timer.Start()
	})

	Convey("Test Run Function", t, func() {
		timer := NewTimer()
		timer.AddJobs(1, func() {
			print(" ")
		}, func() {
			print(" ")
		})
		go timer.Run()
	})

	// time.Sleep(2 * time.Second)
}

func TestTimerStop(t *testing.T) {
	Convey("Test Run Job", t, func() {
		timer := NewTimer()
		timer.AddJob(1, func() {
			print(" ")
		})
		timer.Start()
		// time.Sleep(2 * time.Second)
		timer.Stop()
		// time.Sleep(2 * time.Second)
		timer.Start()
		//time.Sleep(2 * time.Second)
	})

	Convey("Test Run Jobs", t, func() {
		timer := NewTimer()
		timer.AddJobs(1, func() {
			print(" ")
		}, func() {
			print(" ")
		})
		timer.Start()
		time.Sleep(2 * time.Second)
		timer.Stop()
		time.Sleep(2 * time.Second)
		timer.Start()
		// time.Sleep(2 * time.Second)
	})
}

package utils

import (
	"sync"
	"time"
)

// Task is a job.
type Task = func()

// Tasks is a job group.
type Tasks = []func()

// Timer is a Implement of Timer
type Timer struct {
	jobs map[time.Duration]Tasks

	r sync.RWMutex

	stopFlag chan struct{}
}

// NewTimer is return a Timer Instance.
func NewTimer() *Timer {
	return &Timer{
		jobs:     make(map[time.Duration]Tasks),
		stopFlag: make(chan struct{}),
	}
}

// initSecondJobs is Check map key of job is blank?
// if empty then create him.
// else will be nothing.
func (t *Timer) initSecondJobs(seconds time.Duration) {
	if t.jobs[seconds] == nil {
		t.jobs[seconds] = make(Tasks, 0)
	}
}

// AddJob is Add a Job to Jobs.
func (t *Timer) AddJob(seconds time.Duration, job Task) {
	t.initSecondJobs(seconds)
	t.jobs[seconds] = append(t.jobs[seconds], job)
}

// AddJobs is Add a lot of jobs to Jobs.
func (t *Timer) AddJobs(seconds time.Duration, jobs ...Task) {
	t.initSecondJobs(seconds)
	t.jobs[seconds] = append(t.jobs[seconds], jobs...)
}

// Start is Run All Jobs in goroutine.
func (t *Timer) Start() {
	t.forEach(func(time time.Duration, ts Tasks) {
		go t.oneStart(time, ts)
	})
}

// Run All Jobs.
func (t *Timer) Run() {
	t.forEach(func(time time.Duration, ts Tasks) {
		go t.oneStart(time, ts)
	})

	select {}
}

// Stop All Job in Timer
func (t *Timer) Stop() {
	t.stopFlag <- struct{}{}
}

// forEach is foreach all Jobs.
func (t *Timer) forEach(fn func(seconds time.Duration, ts Tasks)) {
	for time, tasks := range t.jobs {
		fn(time, tasks)
	}
}

// oneStart is Start a Job.
func (t *Timer) oneStart(seconds time.Duration, ts Tasks) {
	timer := time.NewTicker(time.Second * seconds)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			for _, task := range ts {
				task()
			}

		case <-t.stopFlag:
			return
		}
	}

}

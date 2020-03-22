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

	stopFlag bool
}

// NewTimer is return a Timer Instance.
func NewTimer() *Timer {
	return &Timer{
		jobs:     make(map[time.Duration]Tasks),
		stopFlag: false,
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
	t.r.Lock()
	t.stopFlag = false
	t.r.Unlock()
	t.forEach(func(time time.Duration, ts Tasks) {
		go t.oneStart(time, ts)
	})
}

// Run All Jobs.
func (t *Timer) Run() {
	t.r.Lock()
	t.stopFlag = false
	t.r.Unlock()
	t.forEach(func(time time.Duration, ts Tasks) {
		go t.oneStart(time, ts)
	})

	select {}
}

// Stop All Job in Timer
func (t *Timer) Stop() {
	t.r.Lock()
	t.stopFlag = true
	defer t.r.Unlock()
}

// forEach is foreach all Jobs.
func (t *Timer) forEach(fn func(seconds time.Duration, ts Tasks)) {
	for time, tasks := range t.jobs {
		fn(time, tasks)
	}
}

// oneStart is Start a Job.
func (t *Timer) oneStart(seconds time.Duration, ts Tasks) {
	timer := time.NewTimer(time.Second * seconds)

	for {
		<-timer.C
		for _, t := range ts {
			t()
		}

		// Lock
		t.r.RLock()
		if t.stopFlag {
			t.r.RUnlock()
			break
		}
		t.r.RUnlock()

		timer.Reset(time.Second * seconds)
	}

	defer timer.Stop()
}
